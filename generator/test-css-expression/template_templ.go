// Code generated by templ - DO NOT EDIT.

package testcssexpression

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"

func className() templ.CSSClass {
	templ_7745c5c3_CSSBuilder := templ.GetBuffer()
	defer templ.ReleaseBuffer(templ_7745c5c3_CSSBuilder)
	templ_7745c5c3_CSSBuilder.WriteString(`background-color:#ffffff;`)
	templ_7745c5c3_CSSBuilder.WriteString(`max-height:calc(100vh - 170px);`)
	templ_7745c5c3_CSSBuilder.WriteString(string(templ.SanitizeCSS(`color`, red)))
	templ_7745c5c3_CSSID := templ.CSSID(`className`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}
