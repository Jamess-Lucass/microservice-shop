using System.ComponentModel.DataAnnotations.Schema;

public record User : Base
{
    [Column(TypeName = "varchar(128)")]
    public string Firstname { get; set; } = default!;

    [Column(TypeName = "varchar(128)")]
    public string Lastname { get; set; } = default!;

    [Column(TypeName = "varchar(320)")]
    public string Email { get; set; } = default!;

    public UserStatus Status { get; set; }
}