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

union LoginIdentifier {
    email: string;
    user_id: string;
}
const LoginRequestExample = #{
  identifier: "test@example.com or test123",
  password: "password",
};
@example(LoginRequestExample)
model LoginRequest {
  identifier: LoginIdentifier;
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

const SuccessResponseExample = #{
  message: "message: Success"
};
@example(SuccessResponseExample)
model SuccessResponse {
  message: string;
}

const ErrorResponseExample = #{
  message: "message: Error"
};
@example(ErrorResponseExample)
@error
model ErrorResponse {
  message: string;
}

@route("/api/v1/users/register")
@post
@summary("Register a new user")
@tag("Users")
@useAuth(NoAuth)
op registerUser(
  @body requestBody: RegisterUserRequest
): RegisterUserResponse | ErrorResponse;

@route("/api/v1/users/me")
@delete
@summary("Delete a user")
@tag("Users")
@useAuth(BearerAuth)
op deleteUser(): SuccessResponse | ErrorResponse;

@route("/api/v1/users/login")
@post
@summary("Login user")
@tag("Users")
@useAuth(BearerAuth)
op login(
  @body requestBody: LoginRequest
): RegisterUserResponse | ErrorResponse;

@route("/api/v1/users/logout")
@post
@summary("Logout user")
@tag("Users")
@useAuth(BearerAuth)
op logout(): SuccessResponse | ErrorResponse;

@route("/api/v1/users/accounts")
@put
@summary("Connect to an external service")
@tag("Users")
@useAuth(BearerAuth)
op connectService(): void | ErrorResponse;

const noteExample = #{
  note_id: "note123",
  title: "My First Note",
  content: "This is the content of my first note.",
  tags: "work, project",
  createdAt: utcDateTime.fromISO("2023-01-01T12:00:00Z"),
  updatedAt: utcDateTime.fromISO("2023-01-01T12:30:00Z")
};
@example(noteExample)
model Note {
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
@post
@summary("Create a new note")
@tag("Notes")
@useAuth(BearerAuth)
op createNote(
  @body requestBody: Note
): Note | ErrorResponse;

@route("/api/v1/notes")
@get
@summary("Get a list of notes")
@tag("Notes")
@useAuth(BearerAuth)
op getNotes(
  @query page?: int32,
  @query pageSize?: int32
): Note | ErrorResponse;

@route("/api/v1/notes/note")
@get
@summary("Get a note")
@tag("Notes")
@useAuth(BearerAuth)
op getNote(
  @query date: utcDateTime;
): Note | ErrorResponse;

@route("/api/v1/notes/list")
@get
@summary("Get a note list")
@tag("Notes")
@useAuth(BearerAuth)
op getNoteList(
  @query from: utcDateTime;
  @query to: utcDateTime;
): Note[] | ErrorResponse;

@route("/api/v1/notes/search")
@get
@summary("Search notes")
@tag("Notes")
@useAuth(BearerAuth)
op searchNotes(
  @query query: string
): SearchNotesResponse | ErrorResponse;

@route("/api/v1/health")
@get
@summary("Check API health")
@tag("Health")
@useAuth(NoAuth)
op healthCheck(): SuccessResponse | ErrorResponse;
