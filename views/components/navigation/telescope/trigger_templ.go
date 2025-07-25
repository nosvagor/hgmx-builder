// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.865
package telescope

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Trigger() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<telescopeTrigger class=\"flex justify-between items-center py-1 pr-6 pl-2 w-20 text-xs rounded-lg transition-colors cursor-pointer outline-none bg-base-600 text-secondary-601 hover:text-primary-401\" _=\"\n\t\t\ton click\n\t\t\t\thalt the event\n\t\t\t\tsend telescope:open\n\t\t\tend\n\t\t\"><svg class=\"mr-1.5 h-4 min-w-4\" viewBox=\"0 0 24 24\" fill=\"currentColor\"><rect x=\"6\" y=\"9\" width=\"8\" height=\"4\" rx=\"1\" transform=\"rotate(-25 10 11)\"></rect> <rect x=\"2\" y=\"12.5\" width=\"3\" height=\"3\" rx=\"1\" transform=\"rotate(-25 3.5 13)\"></rect> <rect x=\"15\" y=\"5\" width=\"4\" height=\"6\" rx=\"2\" transform=\"rotate(-25 17 9)\"></rect> <rect x=\"8.5\" y=\"14\" width=\"1.5\" height=\"6\" rx=\"0.75\" transform=\"rotate(20 9.25 17)\"></rect> <rect x=\"13\" y=\"13\" width=\"1.5\" height=\"7\" rx=\"0.75\" transform=\"rotate(-20 13.75 17)\"></rect></svg> <kbd class=\"text-[0.8em]\" _=\"\n\t\t\t\tinit\n\t\t\t\t\tif navigator.platform.toUpperCase().indexOf(&#39;MAC&#39;) &gt;= 0\n\t\t\t\t\t\tput &#39;⌘K&#39; into me\n\t\t\t\t\telse\n\t\t\t\t\t\tput &#39;Ctrl+K&#39; into me\n\t\t\t\t\tend\n\t\t\t\tend\n\t\t\t\ton click\n\t\t\t\t\thalt the event\n\t\t\t\t\tsend telescope:open\n\t\t\t\tend\n\t\t\t\ton keydown[metaKey and key is &#39;k&#39;] from window\n\t\t\t\t\thalt the event\n\t\t\t\t\tsend telescope:open\n\t\t\t\tend\n\t\t\t\ton keydown[ctrlKey and key is &#39;k&#39;] from window\n\t\t\t\t\thalt the event\n\t\t\t\t\tsend telescope:open\n\t\t\t\tend\n\t\t\t\">Ctrl+K</kbd></telescopeTrigger>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
