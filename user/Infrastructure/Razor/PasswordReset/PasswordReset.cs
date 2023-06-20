namespace User.Infrastructure.Razor.PasswordReset;

public class PasswordResetTemplate : Template
{
    protected override string ViewPath => @"~/Infrastructure/Razor/PasswordReset/PasswordResetView.cshtml";
    public string ResetLink { get; set; } = "";
}