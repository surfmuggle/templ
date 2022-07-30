// Code generated by templ@(devel) DO NOT EDIT.

package testvoid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func render() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = new(bytes.Buffer)
		}
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		ctx, _ = templ.RenderedScriptsFromContext(ctx)
		var_1 := ctx
		ctx = templ.ClearChildren(var_1)
		// Element (void)
		_, err = templBuffer.WriteString("<br>")
		if err != nil {
			return err
		}
		// Element (void)
		_, err = templBuffer.WriteString("<img")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" src=\"https://example.com/image.png\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (void)
		_, err = templBuffer.WriteString("<br>")
		if err != nil {
			return err
		}
		// Element (void)
		_, err = templBuffer.WriteString("<br>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

