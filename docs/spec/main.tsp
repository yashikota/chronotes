import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";
import "@typespec/versioning";

using TypeSpec.Http;
using TypeSpec.Rest;
using TypeSpec.Versioning;

@service({
  title: "Chronotes API",
  description: "Chronotes API specification"
})
@versioned(Versions)
@server("http://localhost:5678", "dev")
@server("https://chronotes.yashikota.com", "prd")
namespace chronotes;

enum Versions {
  v1,
}

// Models
const userExample = #{
  user_id: "test123",
  user_name: "hoge太郎",
  email: "test@example.com",
  createdAt: utcDateTime.fromISO("2024-10-01T19:30:06Z"),
  updatedAt: utcDateTime.fromISO("2024-10-02T03:34:56Z"),
};
@example(userExample)
model User {
  user_id: string;
  user_name: string;
  email: string;
  createdAt: utcDateTime;
  updatedAt: utcDateTime;
  accounts?: Accounts;
}

const RegisterUserRequestExample = #{
  user_id: "test123",
  user_name: "hoge太郎",
  email: "test@example.com",
  password: "password",
};
@example(RegisterUserRequestExample)
model RegisterUserRequest {
  user_id: string;
  user_name: string;
  email: string;
  password: string;
}

const LoginRequestExample = #{
  email: "test@example.com",
  user_id: "test123",
  password: "password",
};
@example(LoginRequestExample)
model LoginRequest {
  user_id?: string;
  email?: string;
  password: string;
}

const RegisterUserResponseExample = #{
  user: userExample,
  token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG4uZG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
};
model RegisterUserResponse {
  user: User;
  token: string;
}

const PromoteUserRequestExample = #{
  password: "password",
};
model PromoteUserRequest {
  password: string;
}

const AccountsExample = #{
  github: "yashikota",
  discord: "9419543265265028",
  slack: "HG824ORHQ250",
  qiita: "yashikota",
  zenn: "yashikota",
  wakatime: "yashikota",
};
model Accounts {
    github_user_id?: string;
    discord_channel_id?: string;
    slack_channel_id?: string;
    qiita_user_id?: string;
    zenn_username?: string;
    wakatime_token?: string;
}

model UpdateUserRequest {
  user_name?: string;
  email?: string;
  password?: string;
  accounts?: Accounts;
}

const SuccessResponseExample = #{
  message: "Success"
};
@example(SuccessResponseExample)
model SuccessResponse {
  message: string;
}

const ErrorResponseExample = #{
  message: "Error"
};
@example(ErrorResponseExample)
@error
model ErrorResponse {
  message: string;
}

@route("/api/v1/auth")
@tag("Auth")
@useAuth(NoAuth)
interface Auth {
    @route("/register")
    @post
    @summary("Register a new user")
    register(
      @body requestBody: RegisterUserRequest
    ): RegisterUserResponse | ErrorResponse;

    @route("/login")
    @post
    @summary("Login user")
    login(
      @body requestBody: LoginRequest
    ): RegisterUserResponse | ErrorResponse;

    @route("/logout")
    @post
    @summary("Logout user")
    @useAuth(BearerAuth)
    logout(): SuccessResponse | ErrorResponse;
}

@route("/api/v1/users/me")
@tag("Users")
@useAuth(BearerAuth)
interface Users {
    @get
    @summary("Get user information")
    get(): User | ErrorResponse;

    @put
    @summary("Update user information")
    update(
      @body requestBody: UpdateUserRequest
    ): User | ErrorResponse;

    @delete
    @summary("Delete a user")
    delete(): SuccessResponse | ErrorResponse;
}

@route("/api/v1/users/promote")
@tag("Users")
@useAuth(BearerAuth)
interface UsersPromote {
    @put
    @summary("Promote a user")
    promote(
      @body requestBody: PromoteUserRequest
    ): SuccessResponse | ErrorResponse;
}

const noteExample = #{
  user_id: "test123",
  note_id: "FHNWFLAQN",
  title: "My First Note",
  content: "This is the content of my first note.",
  tags: "work, project",
  createdAt: utcDateTime.fromISO("2023-01-01T12:00:00Z"),
  updatedAt: utcDateTime.fromISO("2023-01-01T12:30:00Z")
};
@example(noteExample)
model Note {
  user_id: string;
  note_id: string;
  title: string;
  content: string;
  tags: string;
  createdAt: utcDateTime;
  updatedAt: utcDateTime;
}

@example(#{
  title: "Updated Note Title",
  content: "This is the updated content of my note.",
  tags: "go, python"
})
model UpdateNoteRequest {
  title?: string;
  content?: string;
  tags?: string;
}

@example(#{
  notes: #[
    noteExample,
    noteExample
  ]
})
model SearchNotesResponse {
  notes: Note[];
}

@example(#{
  tags: "important, urgent, project"
})
model AddTagsRequest {
  tags: string;
}

@route("/api/v1/notes")
@tag("Notes")
@useAuth(BearerAuth)
interface Notes {
    @route("/list")
    @get
    @summary("Get a list of notes")
    getNotes(
      @query from: utcDateTime;
      @query to: utcDateTime;
    ): Note[] | ErrorResponse;

    @route("/note")
    @get
    @summary("Get a note")
    getNote(
      @query date: utcDateTime;
    ): Note | ErrorResponse;

    @route("/search")
    @post
    @summary("Search notes")
    searchNotes(
      @query query: string
    ): SearchNotesResponse | ErrorResponse;
}

@route("/api/v1/admin")
@tag("Admin")
@useAuth(BearerAuth)
interface Admin {
    @route("/notes")
    @post
    @summary("Add notes")
    addNotes(
      @body requestBody: Note[]
    ): SuccessResponse | ErrorResponse;
}

@route("/api/v1/health")
@get
@summary("Check API health")
@tag("Health")
@useAuth(NoAuth)
op healthCheck(): SuccessResponse | ErrorResponse;
