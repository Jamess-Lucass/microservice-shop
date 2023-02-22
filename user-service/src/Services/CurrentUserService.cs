using System.Security.Claims;

public class CurrentUserService : ICurrentUserService
{
    private readonly IHttpContextAccessor _httpContextAccessor;

    public CurrentUserService(IHttpContextAccessor httpContextAccessor)
    {
        _httpContextAccessor = httpContextAccessor;
    }

    public Guid UserId => _httpContextAccessor.HttpContext?.User.Claims.Where(x => x.Type == ClaimTypes.NameIdentifier).Select(x => Guid.Parse(x.Value)).FirstOrDefault() ?? Guid.Empty;
    public string? FirstName => _httpContextAccessor.HttpContext?.User.Claims.Where(x => x.Type == "firstName").Select(x => x.Value).FirstOrDefault();
    public string? LastName => _httpContextAccessor.HttpContext?.User.Claims.Where(x => x.Type == "LastName").Select(x => x.Value).FirstOrDefault();
    public string? Email => _httpContextAccessor.HttpContext?.User.Claims.Where(x => x.Type == ClaimTypes.Email).Select(x => x.Value).FirstOrDefault();
    public int CustomerId => _httpContextAccessor.HttpContext?.User.Claims.Where(x => x.Type == "customerId").Select(x => Convert.ToInt32(x.Value)).FirstOrDefault() ?? 0;
    public int ResellerId => _httpContextAccessor.HttpContext?.User.Claims.Where(x => x.Type == "resellerId").Select(x => Convert.ToInt32(x.Value)).FirstOrDefault() ?? 0;
    public int OrgId => _httpContextAccessor.HttpContext?.User.Claims.Where(x => x.Type == "orgId").Select(x => Convert.ToInt32(x.Value)).FirstOrDefault() ?? 0;
}