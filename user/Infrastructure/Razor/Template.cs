using Razor.Templating.Core;

namespace User.Infrastructure.Razor;

public abstract class Template
{
    protected abstract string ViewPath { get; }
    public async Task<string> RenderAsync(CancellationToken cancellationToken)
    {
        return await RazorTemplateEngine.RenderAsync(ViewPath, this);
    }
}