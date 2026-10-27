package capmonster

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var testProxy = &Proxy{Type: ProxyHTTP, Address: "1.2.3.4", Port: 8080, Login: "u", Password: "p"}

const proxyJSON = `"proxyType":"http","proxyAddress":"1.2.3.4","proxyPort":8080,"proxyLogin":"u","proxyPassword":"p"`

func TestTaskJSON(t *testing.T) {
	u := "https://example.com"
	cases := []struct {
		name string
		req  Request
		want string // JSON object; key order does not matter
	}{
		{"recaptcha v2 minimal", RecaptchaV2Request{WebsiteURL: u, WebsiteKey: "KEY"},
			`{"type":"RecaptchaV2Task","websiteURL":"https://example.com","websiteKey":"KEY"}`},
		{"recaptcha v2 full", &RecaptchaV2Request{WebsiteURL: u, WebsiteKey: "KEY", Proxy: testProxy, DataSValue: "DS",
			UserAgent: "UA", Cookies: "a=b", IsInvisible: Bool(false)},
			`{"type":"RecaptchaV2Task","websiteURL":"https://example.com","websiteKey":"KEY","recaptchaDataSValue":"DS",
			"userAgent":"UA","cookies":"a=b","isInvisible":false,` + proxyJSON + `}`},
		{"recaptcha v2 enterprise", RecaptchaV2EnterpriseRequest{WebsiteURL: u, WebsiteKey: "KEY",
			EnterprisePayload: map[string]string{"s": "S"}, APIDomain: "www.recaptcha.net", PageAction: "login"},
			`{"type":"RecaptchaV2EnterpriseTask","websiteURL":"https://example.com","websiteKey":"KEY",
			"enterprisePayload":{"s":"S"},"apiDomain":"www.recaptcha.net","pageAction":"login"}`},
		{"recaptcha v3", RecaptchaV3ProxylessRequest{WebsiteURL: u, WebsiteKey: "KEY", MinScore: 0.3, PageAction: "verify", IsEnterprise: Bool(false)},
			`{"type":"RecaptchaV3TaskProxyless","websiteURL":"https://example.com","websiteKey":"KEY","minScore":0.3,
			"pageAction":"verify","isEnterprise":false}`},
		{"recaptcha v3 enterprise", RecaptchaV3EnterpriseRequest{WebsiteURL: u, WebsiteKey: "KEY", MinScore: 0.9},
			`{"type":"RecaptchaV3EnterpriseTask","websiteURL":"https://example.com","websiteKey":"KEY","minScore":0.9}`},
		{"recaptcha click by url", RecaptchaComplexImageTaskRequest{ImageURLs: []string{"https://x/img.jpg"},
			Metadata: RecaptchaMetadata{Grid: Grid3x3, Task: "Click on traffic lights"}},
			`{"type":"ComplexImageTask","class":"recaptcha","imageUrls":["https://x/img.jpg"],
			"metadata":{"Grid":"3x3","Task":"Click on traffic lights"}}`},
		{"recaptcha click by base64", RecaptchaComplexImageTaskRequest{ImagesBase64: []string{"QUJD"},
			Metadata: RecaptchaMetadata{Grid: Grid1x1, TaskDefinition: "/m/015qff"}, UserAgent: "UA", WebsiteURL: u},
			`{"type":"ComplexImageTask","class":"recaptcha","imagesBase64":["QUJD"],
			"metadata":{"Grid":"1x1","TaskDefinition":"/m/015qff"},"userAgent":"UA","websiteURL":"https://example.com"}`},
		{"recognition", RecognitionComplexImageTaskRequest{ImagesBase64: []string{"QUJD"},
			Metadata: RecognitionMetadata{Task: "bills_audio", TaskArgument: "123", PayloadType: PayloadAudio}},
			`{"type":"ComplexImageTask","class":"recognition","imagesBase64":["QUJD"],
			"metadata":{"Task":"bills_audio","TaskArgument":"123","PayloadType":"Audio"}}`},
		{"image to text", ImageToTextRequest{Body: []byte("ABC"), CapMonsterModule: ModuleUniversal,
			RecognizingThreshold: Int(70), Case: Bool(true), Numeric: Bool(true), Math: Bool(false)},
			`{"type":"ImageToTextTask","body":"QUJD","capMonsterModule":"universal","recognizingThreshold":70,
			"case":true,"numeric":1,"math":false}`},
		{"image to text threshold 0 is sent", ImageToTextRequest{Body: []byte("ABC"), RecognizingThreshold: Int(0), Numeric: Bool(false)},
			`{"type":"ImageToTextTask","body":"QUJD","recognizingThreshold":0,"numeric":0}`},
		{"funcaptcha", FuncaptchaRequest{WebsiteURL: u, WebsitePublicKey: "PK", Proxy: testProxy,
			FuncaptchaAPIJSSubdomain: "x.arkoselabs.com", Data: `{"blob":"b"}`},
			`{"type":"FunCaptchaTask","websiteURL":"https://example.com","websitePublicKey":"PK",
			"funcaptchaApiJSSubdomain":"x.arkoselabs.com","data":"{\"blob\":\"b\"}",` + proxyJSON + `}`},
		{"geetest v3", GeetestV3Request{WebsiteURL: u, GT: "GT", Challenge: "CH", GeetestGetLib: `{"static_servers":["x"]}`},
			`{"type":"GeeTestTask","websiteURL":"https://example.com","gt":"GT","version":3,"challenge":"CH",
			"geetestGetLib":"{\"static_servers\":[\"x\"]}"}`},
		{"geetest v4", GeetestV4Request{WebsiteURL: u, GT: "CID", InitParameters: map[string]any{"riskType": "slide"}, Proxy: testProxy, UserAgent: "UA"},
			`{"type":"GeeTestTask","websiteURL":"https://example.com","gt":"CID","version":4,
			"initParameters":{"riskType":"slide"},"userAgent":"UA",` + proxyJSON + `}`},
		{"turnstile standard", TurnstileRequest{WebsiteURL: u, WebsiteKey: "0xKEY", PageAction: "login", Data: "cdata"},
			`{"type":"TurnstileTask","websiteURL":"https://example.com","websiteKey":"0xKEY","pageAction":"login","data":"cdata"}`},
		{"turnstile challenge token", TurnstileChallengeTokenRequest{WebsiteURL: u, WebsiteKey: "0xKEY", PageAction: "managed",
			UserAgent: "UA", Data: "cdata", PageData: "pdata", APIJSURL: "https://x/api.js"},
			`{"type":"TurnstileTask","websiteURL":"https://example.com","websiteKey":"0xKEY","cloudflareTaskType":"token",
			"pageAction":"managed","userAgent":"UA","data":"cdata","pageData":"pdata","apiJsUrl":"https://x/api.js"}`},
		{"turnstile cf_clearance", TurnstileChallengeCookieRequest{WebsiteURL: u, WebsiteKey: "xxxxxx", HTMLPageBase64: "SFRNTA==",
			UserAgent: "UA", Proxy: testProxy},
			`{"type":"TurnstileTask","websiteURL":"https://example.com","websiteKey":"xxxxxx","cloudflareTaskType":"cf_clearance",
			"userAgent":"UA","htmlPageBase64":"SFRNTA==",` + proxyJSON + `}`},
		{"turnstile wait room", TurnstileWaitRoomRequest{WebsiteURL: u, WebsiteKey: "xxxxxx", HTMLPageBase64: "SFRNTA==",
			UserAgent: "UA", Proxy: testProxy},
			`{"type":"TurnstileTask","websiteURL":"https://example.com","websiteKey":"xxxxxx","cloudflareTaskType":"wait_room",
			"userAgent":"UA","htmlPageBase64":"SFRNTA==",` + proxyJSON + `}`},
		{"amazon visible", AmazonWafVisibleRequest{WebsiteURL: u, WebsiteKey: "KEY", CaptchaScript: "https://x/jsapi.js",
			Proxy: testProxy, UserAgent: "UA", CookieSolution: Bool(true)},
			`{"type":"AmazonTask","websiteURL":"https://example.com","websiteKey":"KEY","captchaScript":"https://x/jsapi.js",
			"userAgent":"UA","cookieSolution":true,` + proxyJSON + `}`},
		{"amazon challenge and captcha", AmazonWafChallengeRequest{WebsiteURL: u, ChallengeScript: "https://x/challenge.js",
			WebsiteKey: "K", Context: "C", IV: "IV", CaptchaScript: "https://x/captcha.js", CookieSolution: Bool(false)},
			`{"type":"AmazonTask","websiteURL":"https://example.com","websiteKey":"K","captchaScript":"https://x/captcha.js",
			"challengeScript":"https://x/challenge.js","context":"C","iv":"IV","cookieSolution":false}`},
		{"amazon challenge only sends empty context and iv", AmazonWafChallengeOnlyRequest{WebsiteURL: u,
			ChallengeScript: "https://x/challenge.js", CookieSolution: Bool(true)},
			`{"type":"AmazonTask","websiteURL":"https://example.com","challengeScript":"https://x/challenge.js",
			"context":"","iv":"","cookieSolution":true}`},
		{"binance", BinanceTaskRequest{WebsiteURL: u, WebsiteKey: "login", ValidateID: "vid", UserAgent: "UA"},
			`{"type":"BinanceTask","websiteURL":"https://example.com","websiteKey":"login","validateId":"vid","userAgent":"UA"}`},
		{"mtcaptcha", MTCaptchaRequest{WebsiteURL: u, WebsiteKey: "MTPublic-x", PageAction: "login", IsInvisible: Bool(true)},
			`{"type":"MTCaptchaTask","websiteURL":"https://example.com","websiteKey":"MTPublic-x","pageAction":"login","isInvisible":true}`},
		{"yidun enterprise", YidunRequest{WebsiteURL: u, WebsiteKey: "SK", YidunGetLib: "https://x/load.js",
			YidunAPIServerSubdomain: "c.example.com", Challenge: "CH", Hcg: "H", Hct: 9007199254740993},
			`{"type":"YidunTask","websiteURL":"https://example.com","websiteKey":"SK","yidunGetLib":"https://x/load.js",
			"yidunApiServerSubdomain":"c.example.com","challenge":"CH","hcg":"H","hct":9007199254740993}`},
		{"prosopo", ProsopoTaskRequest{WebsiteURL: u, WebsiteKey: "SK"},
			`{"type":"ProsopoTask","websiteURL":"https://example.com","websiteKey":"SK"}`},
		{"datadome", DataDomeCustomTaskRequest{WebsiteURL: u, Proxy: testProxy, UserAgent: "UA", Metadata: DataDomeMetadata{
			CaptchaURL: "https://geo.captcha-delivery.com/captcha/?t=fe", DatadomeCookie: "datadome=abc", DatadomeVersion: DataDomeVersionNew}},
			`{"type":"CustomTask","class":"DataDome","websiteURL":"https://example.com",
			"metadata":{"captchaUrl":"https://geo.captcha-delivery.com/captcha/?t=fe","datadomeCookie":"datadome=abc","datadomeVersion":"new"},
			"userAgent":"UA",` + proxyJSON + `}`},
		{"imperva", ImpervaCustomTaskRequest{WebsiteURL: u, Proxy: testProxy, Metadata: ImpervaMetadata{
			IncapsulaScriptURL: "_Incapsula_Resource?x=1", IncapsulaCookies: "incap_ses=1", Reese84URLEndpoint: "/endpoint?d=site.com"}},
			`{"type":"CustomTask","class":"Imperva","websiteURL":"https://example.com","metadata":{"incapsulaScriptUrl":"_Incapsula_Resource?x=1",
			"incapsulaCookies":"incap_ses=1","reese84UrlEndpoint":"/endpoint?d=site.com"},` + proxyJSON + `}`},
		{"tspd", TspdCustomTaskRequest{WebsiteURL: u, Proxy: testProxy, Metadata: TspdMetadata{TspdCookie: "TS1=2;", HTMLPageBase64: "SFRNTA=="}},
			`{"type":"CustomTask","class":"tspd","websiteURL":"https://example.com",
			"metadata":{"tspdCookie":"TS1=2;","htmlPageBase64":"SFRNTA=="},` + proxyJSON + `}`},
		{"hunt fingerprint", HuntCustomTaskRequest{WebsiteURL: u, Proxy: testProxy, Metadata: HuntMetadata{APIGetLib: "https://x/api.js"}},
			`{"type":"CustomTask","class":"HUNT","websiteURL":"https://example.com","metadata":{"apiGetLib":"https://x/api.js"},` + proxyJSON + `}`},
		{"hunt widget url", HuntCustomTaskRequest{WebsiteURL: u, Proxy: testProxy, UserAgent: "UA",
			Metadata: HuntMetadata{APIGetLib: "https://x/api.js", WidgetURL: "https://c/widget?hash=1"}},
			`{"type":"CustomTask","class":"HUNT","websiteURL":"https://example.com",
			"metadata":{"apiGetLib":"https://x/api.js","widgetUrl":"https://c/widget?hash=1"},"userAgent":"UA",` + proxyJSON + `}`},
		{"basilisk", BasiliskCustomTaskRequest{WebsiteURL: u, WebsiteKey: "KEY", UserAgent: "UA"},
			`{"type":"CustomTask","class":"Basilisk","websiteURL":"https://example.com","websiteKey":"KEY","userAgent":"UA"}`},
		{"tendi minimal has no metadata", TenDiCustomTaskRequest{WebsiteURL: u, WebsiteKey: "189123456"},
			`{"type":"CustomTask","class":"TenDI","websiteURL":"https://example.com","websiteKey":"189123456"}`},
		{"tendi with metadata", TenDiCustomTaskRequest{WebsiteURL: u, WebsiteKey: "189123456",
			Metadata: &TenDiMetadata{CaptchaURL: "https://x/TCaptcha.js", AidEncrypted: "AID"}},
			`{"type":"CustomTask","class":"TenDI","websiteURL":"https://example.com","websiteKey":"189123456",
			"metadata":{"captchaUrl":"https://x/TCaptcha.js","aidEncrypted":"AID"}}`},
		{"alibaba", AlibabaCustomTaskRequest{WebsiteURL: u, Metadata: AlibabaMetadata{SceneID: "S", Prefix: "P", Region: "sgp",
			UserCertifyID: "CID", CookieRequired: Bool(true)}},
			`{"type":"CustomTask","class":"alibaba","websiteURL":"https://example.com",
			"metadata":{"sceneId":"S","prefix":"P","region":"sgp","UserCertifyId":"CID","cookieRequired":true}}`},
		{"friendly", FriendlyCustomTaskRequest{WebsiteURL: u, WebsiteKey: "FCKEY", Metadata: FriendlyMetadata{APIGetLib: "https://x/site.min.js"}},
			`{"type":"CustomTask","class":"friendly","websiteURL":"https://example.com","websiteKey":"FCKEY",
			"metadata":{"apiGetLib":"https://x/site.min.js"}}`},
		{"altcha with empty website key", AltchaCustomTaskRequest{WebsiteURL: u,
			Metadata: AltchaMetadata{Challenge: "c", Iterations: "1", Salt: "s", Signature: "g"}},
			`{"type":"CustomTask","class":"altcha","websiteURL":"https://example.com","websiteKey":"",
			"metadata":{"challenge":"c","iterations":"1","salt":"s","signature":"g"}}`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			task, err := tc.req.Task()
			if err != nil {
				t.Fatalf("Task(): %v", err)
			}
			assertJSONEqual(t, task, tc.want)
			if tc.req.Timeouts().Timeout <= 0 {
				t.Error("no timeout in the polling schedule")
			}
		})
	}
}

func TestValidation(t *testing.T) {
	u := "https://example.com"
	cases := []struct {
		name  string
		req   Request
		field string
	}{
		{"empty websiteURL", RecaptchaV2Request{WebsiteKey: "KEY"}, "websiteURL"},
		{"empty websiteKey", TurnstileRequest{WebsiteURL: u}, "websiteKey"},
		{"minScore too high", RecaptchaV3ProxylessRequest{WebsiteURL: u, WebsiteKey: "K", MinScore: 1.1}, "minScore"},
		{"minScore too low (enterprise)", RecaptchaV3EnterpriseRequest{WebsiteURL: u, WebsiteKey: "K", MinScore: 0.05}, "minScore"},
		{"threshold out of range", ImageToTextRequest{Body: []byte("x"), RecognizingThreshold: Int(101)}, "recognizingThreshold"},
		{"empty image body", ImageToTextRequest{}, "body"},
		{"data uri body", ImageToTextRequest{Body: []byte("data:image/png;base64,QUJD")}, "body"},
		{"geetest v3 empty challenge", GeetestV3Request{WebsiteURL: u, GT: "GT"}, "challenge"},
		{"turnstile token without pageData", TurnstileChallengeTokenRequest{WebsiteURL: u, WebsiteKey: "K", PageAction: "managed",
			UserAgent: "UA", Data: "d"}, "pageData"},
		{"turnstile cookie without proxy", TurnstileChallengeCookieRequest{WebsiteURL: u, WebsiteKey: "K", HTMLPageBase64: "H",
			UserAgent: "UA"}, "proxy"},
		{"amazon visible without script", AmazonWafVisibleRequest{WebsiteURL: u, WebsiteKey: "K"}, "captchaScript"},
		{"amazon option 2 without iv", AmazonWafChallengeRequest{WebsiteURL: u, ChallengeScript: "S", WebsiteKey: "K", Context: "C"}, "iv"},
		{"recaptcha click without task", RecaptchaComplexImageTaskRequest{ImageURLs: []string{"u"}, Metadata: RecaptchaMetadata{Grid: Grid3x3}}, "metadata"},
		{"recaptcha click without images", RecaptchaComplexImageTaskRequest{Metadata: RecaptchaMetadata{Grid: Grid3x3, Task: "T"}}, "imageUrls"},
		{"recaptcha click two images", RecaptchaComplexImageTaskRequest{ImageURLs: []string{"a", "b"},
			Metadata: RecaptchaMetadata{Grid: Grid3x3, Task: "T"}}, "imageUrls"},
		{"recognition without images", RecognitionComplexImageTaskRequest{Metadata: RecognitionMetadata{Task: "shein"}}, "imagesBase64"},
		{"recognition without task", RecognitionComplexImageTaskRequest{ImagesBase64: []string{"QUJD"}}, "metadata.Task"},
		{"funcaptcha without proxy", FuncaptchaRequest{WebsiteURL: u, WebsitePublicKey: "PK"}, "proxy"},
		{"datadome empty cookie", DataDomeCustomTaskRequest{WebsiteURL: u, Proxy: testProxy, Metadata: DataDomeMetadata{CaptchaURL: "U"}},
			"metadata.datadomeCookie"},
		{"hunt data and widget url", HuntCustomTaskRequest{WebsiteURL: u, Proxy: testProxy,
			Metadata: HuntMetadata{APIGetLib: "L", Data: "D", WidgetURL: "W"}}, "metadata"},
		{"tspd bad proxy login", TspdCustomTaskRequest{WebsiteURL: u, Metadata: TspdMetadata{TspdCookie: "C", HTMLPageBase64: "H"},
			Proxy: &Proxy{Type: ProxyHTTP, Address: "h", Port: 1, Login: "a|b"}}, "proxyLogin"},
		{"altcha empty salt", AltchaCustomTaskRequest{WebsiteURL: u, Metadata: AltchaMetadata{Challenge: "c", Iterations: "1", Signature: "g"}},
			"metadata.salt"},
		{"alibaba empty prefix", AlibabaCustomTaskRequest{WebsiteURL: u, Metadata: AlibabaMetadata{SceneID: "S"}}, "metadata.prefix"},
		{"binance empty validateId", BinanceTaskRequest{WebsiteURL: u, WebsiteKey: "login"}, "validateId"},
		{"proxy bad port", RecaptchaV2Request{WebsiteURL: u, WebsiteKey: "K", Proxy: &Proxy{Type: ProxySOCKS5, Address: "h", Port: 70000}}, "proxyPort"},
		{"proxy bad type", RecaptchaV2Request{WebsiteURL: u, WebsiteKey: "K", Proxy: &Proxy{Type: "ftp", Address: "h", Port: 1}}, "proxyType"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := tc.req.Task()
			var ve *ValidationError
			if !errors.As(err, &ve) {
				t.Fatalf("want *ValidationError, got %v", err)
			}
			if !strings.HasPrefix(ve.Field, tc.field) {
				t.Errorf("field = %q, want prefix %q (%v)", ve.Field, tc.field, err)
			}
		})
	}
}

func TestImageToTextFromFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "captcha.png")
	if err := os.WriteFile(path, []byte("ABC"), 0o600); err != nil {
		t.Fatal(err)
	}
	req, err := ImageToTextFromFile(path)
	if err != nil {
		t.Fatal(err)
	}
	req.CapMonsterModule = ModuleYandex
	task, err := req.Task()
	if err != nil {
		t.Fatal(err)
	}
	assertJSONEqual(t, task, `{"type":"ImageToTextTask","body":"QUJD","capMonsterModule":"yandex"}`)
}

func TestProxyPasswordIsMaskedInLogs(t *testing.T) {
	req := RecaptchaV2Request{WebsiteURL: "u", WebsiteKey: "k", Proxy: &Proxy{Type: ProxyHTTP, Address: "h", Port: 1, Login: "u", Password: "SECRET"}}
	for _, verb := range []string{"%v", "%+v", "%#v", "%s"} {
		if out := fmt.Sprintf(verb, req); strings.Contains(out, "SECRET") {
			t.Errorf("%s leaks the password: %s", verb, out)
		}
	}
}

func assertJSONEqual(t *testing.T, got any, want string) {
	t.Helper()
	gotJSON, err := json.Marshal(got)
	if err != nil {
		t.Fatal(err)
	}
	var g, w any
	if err := json.Unmarshal(gotJSON, &g); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal([]byte(want), &w); err != nil {
		t.Fatalf("bad expected JSON: %v", err)
	}
	gn, _ := json.Marshal(g)
	wn, _ := json.Marshal(w)
	if string(gn) != string(wn) {
		t.Errorf("task JSON mismatch\n got: %s\nwant: %s", gn, wn)
	}
}
