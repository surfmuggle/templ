// Code generated by templ - DO NOT EDIT.

package testscriptinline

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"

func withParameters(a string, b string, c int) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_withParameters_1056`,
		Function: `function __templ_withParameters_1056(a, b, c){console.log(a, b, c);
}`,
		Call:       templ.SafeScript(`__templ_withParameters_1056`, a, b, c),
		CallInline: templ.SafeScriptInline(`__templ_withParameters_1056`, a, b, c),
	}
}

func withoutParameters() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_withoutParameters_6bbf`,
		Function: `function __templ_withoutParameters_6bbf(){alert("hello");
}`,
		Call:       templ.SafeScript(`__templ_withoutParameters_6bbf`),
		CallInline: templ.SafeScriptInline(`__templ_withoutParameters_6bbf`),
	}
}

func InlineJavascript(a string) templ.Component {
	return templ.GeneratedTemplate(func(templ_7745c5c3_Input templ.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer, templ_7745c5c3_Release := templ.WriterToBuffer(templ_7745c5c3_W)
		defer templ_7745c5c3_Release()
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = withoutParameters().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = withParameters(a, "test", 123).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = withoutParameters().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = withParameters(a, "test", 123).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
