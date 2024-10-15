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
  accounts: AccountsExample
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
  github_user_id: "yashikota",
  discord_channel_id: "9419543265265028",
  slack_channel_id: "HG824ORHQ250",
  qiita_user_id: "yashikota",
  zenn_username: "yashikota",
  wakatime_token: "yashikota",
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
  @get
  @summary("Get a list of notes")
  @doc("fieldsは `note_id`, `user_id`, `title`, `content`, `length`, `tags`が指定可能。複数指定する場合は空白なしのコンマで指定 (例) `title,tags`")
  getNotes(
    @query from: utcDateTime;
    @query to: utcDateTime;
    @query fields: string;
  ): Note[] | ErrorResponse;

  @put
  @summary("Update a notes")
  updateNote(
    @body requestBody: Note[]
  ): Note | ErrorResponse;

  @delete
  @summary("Delete a notes")
  deleteNote(
    @body requestBody: Note[]
  ): SuccessResponse | ErrorResponse;

  @route("/search")
  @post
  @summary("Search notes")
  searchNotes(
    @query query: string
  ): SearchNotesResponse | ErrorResponse;
}

const UploadImageResponseExample = #{
  converted: "img/01JA0QTJB2MHTBNJV3SJ9KYN9J.jpg",
  original: "img/01JA0QTJB2MHTBNJV3SJ9KYN9J_original.jpg",
};
model UploadImageResponse {
  "converted": string,
  "original": string,
}

@route("/api/v1/images")
@tag("Images")
@useAuth(BearerAuth)
interface Images {
  @post
  @summary("Upload an image")
  @doc("対応画像形式はpngとjpegのみ。ファイルサイズは10MBまで。")
  uploadImage(
    @header contentType: "multipart/form-data", image: bytes
  ): UploadImageResponse | ErrorResponse;

  @delete
  @summary("Delete an image")
  @doc("`image_path`には画像のIDを指定して削除")
  deleteImage(
    @query image_id: string
  ): SuccessResponse | ErrorResponse;
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

  @route("/notes")
  @delete
  @summary("Delete notes")
  deleteNotes(
    @query user_id: string
  ): SuccessResponse | ErrorResponse;
}

@route("/api/v1/health")
@get
@summary("Check API health")
@tag("Health")
@useAuth(NoAuth)
op healthCheck(): SuccessResponse | ErrorResponse;
