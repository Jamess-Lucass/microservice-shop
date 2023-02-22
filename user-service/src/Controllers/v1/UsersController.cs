
using System.Runtime.Serialization.Formatters.Binary;
using System.Text;
using System.Text.Json;
using Asp.Versioning;
using AutoMapper;
using FluentValidation;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.OData.Query;
using static Microsoft.AspNetCore.Http.StatusCodes;

namespace API.Controllers.v1;

[ApiVersion(1.0)]
public class UsersController : BaseODataController
{
    private readonly ApplicationDbContext _context;
    private readonly IUserService _userService;
    private readonly ICurrentUserService _user;
    private readonly ILogger<UsersController> _logger;
    private readonly IMapper _mapper;

    public UsersController(
        ILogger<UsersController> logger,
        IUserService userService,
        ICurrentUserService user,
        ApplicationDbContext context,
        IMapper mapper)
    {
        _logger = logger;
        _userService = userService;
        _user = user;
        _context = context;
        _mapper = mapper;
    }

    // GET: /api/v1/users
    [HttpGet]
    [ProducesResponseType(typeof(IEnumerable<UserDto>), Status200OK)]
    [EnableQuery]
    public ActionResult Get()
    {
        return Ok(_userService.GetAllUsers());
    }
}