package capmonster

import (
	"bytes"
	"fmt"
	"os"
)

// Recognition modules of ImageToTextTask (ImageToTextRequest.CapMonsterModule), as listed in
// https://docs.capmonster.cloud/docs/captchas/ImageToText/module-name.txt. The list grows: any other module
// name can be passed as a plain string.
const (
	ModuleAmazon                  = "amazon"
	ModuleApple                   = "apple"
	ModuleBotdetect               = "botdetect"
	ModuleFacebook                = "facebook"
	ModuleGmx                     = "gmx"
	ModuleGoogle                  = "google"
	ModuleHotmail                 = "hotmail"
	ModuleMailru                  = "mailru"
	ModuleOkeng                   = "okeng"
	ModuleOkrus                   = "okrus"
	ModulePartiallyblur           = "partiallyblur"
	ModuleRamblerrus              = "ramblerrus"
	ModuleRamblerrusnew           = "ramblerrusnew"
	ModuleSolvemedia              = "solvemedia"
	ModuleSteam                   = "steam"
	ModuleVk                      = "vk"
	ModuleWhatsapp                = "whatsapp"
	ModuleWikimediaPaddleEnsemble = "wikimedia_paddle_ensemble"
	ModuleYandex                  = "yandex"
	ModuleYandexnew               = "yandexnew" // two-word Yandex captcha
	ModuleYandexwave              = "yandexwave"
	ModuleYandexwavelatin         = "yandexwavelatin" // Latin Yandex captcha
	ModuleCaptchaMath             = "captcha_math"    // do not combine with Math
	ModuleSat                     = "sat"
	ModuleBlsText                 = "bls_text"
	ModuleUniversal               = "universal" // all other text captchas
)

// ImageToTextRequest is a text-on-image captcha (ImageToTextTask). Solution: text.
type ImageToTextRequest struct {
	Body                 []byte `json:"body"` // raw image bytes (PNG/JPEG/...), sent base64-encoded
	CapMonsterModule     string `json:"capMonsterModule,omitempty"`
	RecognizingThreshold *int   `json:"recognizingThreshold,omitempty"` // minimum confidence, 0..100
	Case                 *bool  `json:"case,omitempty"`                 // case-sensitive text
	Numeric              *bool  `json:"-"`                              // only digits (sent as 0/1)
	Math                 *bool  `json:"math,omitempty"`                 // calculate the expression on the image
}

// ImageToTextFromFile reads an image file into ImageToTextRequest.Body.
func ImageToTextFromFile(path string) (ImageToTextRequest, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return ImageToTextRequest{}, fmt.Errorf("capmonster: read image: %w", err)
	}
	return ImageToTextRequest{Body: body}, nil
}

func (r ImageToTextRequest) Timeouts() Timeouts { return complexImageTimeouts() }

func (r ImageToTextRequest) Task() (map[string]any, error) {
	if len(r.Body) == 0 {
		return nil, &ValidationError{Field: "body", Reason: "must not be empty"}
	}
	if bytes.HasPrefix(r.Body, []byte("data:")) {
		return nil, &ValidationError{Field: "body", Reason: "must be raw image bytes, not a data: URI"}
	}
	if t := r.RecognizingThreshold; t != nil && (*t < 0 || *t > 100) {
		return nil, &ValidationError{Field: "recognizingThreshold", Reason: fmt.Sprintf("must be between 0 and 100, got %d", *t)}
	}
	task, err := buildTask("ImageToTextTask", "", r, nil)
	if err == nil && r.Numeric != nil {
		task["numeric"] = 0
		if *r.Numeric {
			task["numeric"] = 1
		}
	}
	return task, err
}

// RecaptchaGrid is the image grid of a reCAPTCHA click challenge.
type RecaptchaGrid string

const (
	Grid3x3 RecaptchaGrid = "3x3"
	Grid4x4 RecaptchaGrid = "4x4"
	Grid1x1 RecaptchaGrid = "1x1"
)

// RecaptchaMetadata needs Grid and at least one of Task / TaskDefinition.
type RecaptchaMetadata struct {
	Grid RecaptchaGrid `json:"Grid"`
	Task string        `json:"Task,omitempty"` // task name in English, e.g. "Click on traffic lights"
	// TaskDefinition is the technical task id, e.g. "/m/015qff", from the /recaptcha/<api>/reload or
	// /userverify response.
	TaskDefinition string `json:"TaskDefinition,omitempty"`
}

// RecaptchaComplexImageTaskRequest is the reCAPTCHA v2 image click task (ComplexImageTask, class recaptcha).
// No proxy is needed. Pass the full 4x4 / 3x3 image or a new 1x1 tile, by URL or as base64 (exactly one image).
type RecaptchaComplexImageTaskRequest struct {
	ImageURLs []string `json:"imageUrls,omitempty"`
	// ImagesBase64 are base64 images without the data:image/...;base64, prefix.
	ImagesBase64 []string          `json:"imagesBase64,omitempty"`
	Metadata     RecaptchaMetadata `json:"metadata"`
	UserAgent    string            `json:"userAgent,omitempty"` // used to load images passed as URLs
	WebsiteURL   string            `json:"websiteURL,omitempty"`
}

func (r RecaptchaComplexImageTaskRequest) Timeouts() Timeouts { return complexImageTimeouts() }

func (r RecaptchaComplexImageTaskRequest) Task() (map[string]any, error) {
	if r.Metadata.Grid == "" {
		return nil, &ValidationError{Field: "metadata.Grid", Reason: "must not be empty"}
	}
	if r.Metadata.Task == "" && r.Metadata.TaskDefinition == "" {
		return nil, &ValidationError{Field: "metadata", Reason: "either Task or TaskDefinition must be set"}
	}
	if r.ImageURLs == nil && r.ImagesBase64 == nil {
		return nil, &ValidationError{Field: "imageUrls / imagesBase64", Reason: "at least one must be set"}
	}
	if err := checkImageList("imageUrls", r.ImageURLs, 1); err != nil {
		return nil, err
	}
	if err := checkImageList("imagesBase64", r.ImagesBase64, 1); err != nil {
		return nil, err
	}
	return buildTask("ComplexImageTask", "recaptcha", r, nil)
}

// PayloadAudio is the RecognitionMetadata.PayloadType for audio modules (ImagesBase64 then holds base64 audio).
const PayloadAudio = "Audio"

// RecognitionMetadata selects the recognition module.
type RecognitionMetadata struct {
	// Task is the module: shein (click coordinates), MathSum, dli (text), bls_3x3 (grid, needs TaskArgument),
	// baidu, oocl_rotate_new, oocl_rotate_double_new, betpunch_3x3_rotate (rotation angles),
	// portugal_text_find_icon (needs TaskArgument), bills_audio (PayloadType = PayloadAudio).
	Task         string `json:"Task"`
	TaskArgument string `json:"TaskArgument,omitempty"` // e.g. the number to find for bls_3x3 ("123")
	PayloadType  string `json:"PayloadType,omitempty"`  // PayloadAudio for audio modules
}

// RecognitionComplexImageTaskRequest runs a recognition module of ComplexImageTask (class recognition).
// No proxy is needed; most modules take exactly one image.
type RecognitionComplexImageTaskRequest struct {
	ImagesBase64 []string            `json:"imagesBase64"` // base64 without a data: prefix; at least one
	Metadata     RecognitionMetadata `json:"metadata"`
	UserAgent    string              `json:"userAgent,omitempty"`
}

func (r RecognitionComplexImageTaskRequest) Timeouts() Timeouts { return complexImageTimeouts() }

func (r RecognitionComplexImageTaskRequest) Task() (map[string]any, error) {
	if err := required("metadata.Task", r.Metadata.Task); err != nil {
		return nil, err
	}
	if len(r.ImagesBase64) == 0 {
		return nil, &ValidationError{Field: "imagesBase64", Reason: "at least one image expected"}
	}
	if err := checkImageList("imagesBase64", r.ImagesBase64, len(r.ImagesBase64)); err != nil {
		return nil, err
	}
	return buildTask("ComplexImageTask", "recognition", r, nil)
}

func complexImageTimeouts() Timeouts { return timeouts(0.35, 0.2, 10) }

// checkImageList: when given (non-nil), the list must hold 1..max non-empty strings.
func checkImageList(field string, list []string, max int) error {
	if list == nil {
		return nil
	}
	if len(list) == 0 {
		return &ValidationError{Field: field, Reason: "at least one entry expected, got 0"}
	}
	if len(list) > max {
		return &ValidationError{Field: field, Reason: fmt.Sprintf("at most %d entries allowed, got %d", max, len(list))}
	}
	for i, item := range list {
		if item == "" {
			return &ValidationError{Field: fmt.Sprintf("%s[%d]", field, i), Reason: "must not be empty"}
		}
	}
	return nil
}
