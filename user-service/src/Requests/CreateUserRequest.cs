public record CreateUserRequest
{
    public CreateUserRequest(string firstName)
    {
        FirstName = firstName;
    }

    public string FirstName { get; set; }
};