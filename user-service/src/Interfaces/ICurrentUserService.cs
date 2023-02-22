public interface ICurrentUserService
{
    Guid UserId { get; }
    int CustomerId { get; }
    int ResellerId { get; }
    int OrgId { get; }
    string? FirstName { get; }
    string? LastName { get; }
    string? Email { get; }
}