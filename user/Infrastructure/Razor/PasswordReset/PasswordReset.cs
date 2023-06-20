namespace User.Infrastructure.Razor.PasswordReset;

public class PasswordReset : ITemplate
{
    string ITemplate.ViewPath { get; } = @"~/Infrastructure/Razor/PasswordReset/PasswordResetView.cshtml";
    public string ResetLink { get; set; } = "";
}