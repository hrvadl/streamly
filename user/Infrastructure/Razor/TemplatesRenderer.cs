using Razor.Templating.Core;

namespace User.Infrastructure.Razor;

public static class TemplatesRenderer
{
    public static async Task<string> RenderAsync<TTemplate>(TTemplate template, CancellationToken cancellationToken) where TTemplate : ITemplate
    {
        return await RazorTemplateEngine.RenderAsync(template.ViewPath, template);
    }
}