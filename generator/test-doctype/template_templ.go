// Code generated by templ@(devel) DO NOT EDIT.

package testdoctype

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Layout(title, content string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = new(bytes.Buffer)
		}
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		ctx, _ = templ.RenderedScriptsFromContext(ctx)
		var_1 := ctx
		ctx = templ.ClearChildren(var_1)
		// DocType
		_, err = templBuffer.WriteString(`<!doctype html>`)
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<html")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" lang=\"en\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<head>")
		if err != nil {
			return err
		}
		// Element (void)
		_, err = templBuffer.WriteString("<meta")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" charset=\"UTF-8\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (void)
		_, err = templBuffer.WriteString("<meta")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" http-equiv=\"X-UA-Compatible\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" content=\"IE=edge\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (void)
		_, err = templBuffer.WriteString("<meta")
		if err != nil {
			return err
		}
		// Element Attributes
		_, err = templBuffer.WriteString(" name=\"viewport\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" content=\"width=device-width, initial-scale=1.0\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(">")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<title>")
		if err != nil {
			return err
		}
		// StringExpression
		_, err = templBuffer.WriteString(templ.EscapeString(title))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</title>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</head>")
		if err != nil {
			return err
		}
		// Element (standard)
		_, err = templBuffer.WriteString("<body>")
		if err != nil {
			return err
		}
		// StringExpression
		_, err = templBuffer.WriteString(templ.EscapeString(content))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</body>")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</html>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

