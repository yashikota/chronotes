import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";
import "@typespec/versioning";

using TypeSpec.Http;
using TypeSpec.Rest;
using TypeSpec.Versioning;

@service({
  title: "Chronotes API",
  description: "Chronotes api specification"
})
@versioned(Versions)
@server("http://localhost:5678", "dev")
@server("https://chronotes.yashikota.com", "prd")
namespace chronotes;

enum Versions {
    v1,
}

model User {
  id: string;
  username: string;
  email: string;
  createdAt: utcDateTime;
}

model RegisterUserRequest {
  username: string;
  email: string;
  password: string;
}

model LoginRequest {
  email: string;
  password: string;
}

model RegisterUserResponse {
  user: User;
  token: string;
}

@error
model ErrorResponse {
  message: string;
}

@route("/api/v1/users/register")
@post
@summary("Register a new user")
@tag("Users")
op registerUser(
  @body requestBody: RegisterUserRequest
): RegisterUserResponse | ErrorResponse;

@route("/api/v1/users/{userId}")
@delete
@summary("Delete a user")
@tag("Users")
op deleteUser(
  @path userId: string
): void | ErrorResponse;

@route("/api/v1/users/login")
@post
@summary("Login user")
@tag("Users")
op login(
  @body requestBody: LoginRequest
): RegisterUserResponse | ErrorResponse;

@route("/api/v1/users/logout")
@post
@summary("Logout user")
@tag("Users")
op logout(): void | ErrorResponse;

model Note {
  id: string;
  title: string;
  content: string;
  tags: string[];
  images: Image[];
  isDraft: boolean;
  createdAt: utcDateTime;
  updatedAt: utcDateTime;
}

model CreateNoteRequest {
  title: string;
  content: string;
  tags?: string[];
  images?: Image[];
  isDraft?: boolean;
}

model UpdateNoteRequest {
  title?: string;
  content?: string;
  tags?: string[];
  images?: Image[];
  isDraft?: boolean;
}

model GetNotesResponse {
  notes: Note[];
  totalCount: int32;
}

model SearchNotesResponse {
  notes: Note[];
}

model AddTagsRequest {
  tags: string[];
}

model UploadImageRequest {
  image: Image;
}

model Image {
  id: string;
  url: string;
}

@route("/api/v1/notes")
@post
@summary("Create a new note")
@tag("Notes")
op createNote(
  @body requestBody: CreateNoteRequest
): Note | ErrorResponse;

@route("/api/v1/notes/{noteId}")
@put
@summary("Update an existing note")
@tag("Notes")
op updateNote(
  @path noteId: string,
  @body requestBody: UpdateNoteRequest
): Note | ErrorResponse;

@route("/api/v1/notes/{noteId}")
@delete
@summary("Delete a note")
@tag("Notes")
op deleteNote(
  @path noteId: string
): void | ErrorResponse;

@route("/api/v1/notes")
@get
@summary("Get a list of notes")
@tag("Notes")
op getNotes(
  @query page?: int32,
  @query pageSize?: int32
): GetNotesResponse | ErrorResponse;

@route("/api/v1/notes/{noteId}")
@get
@summary("Get a specific note")
@tag("Notes")
op getNote(
  @path noteId: string
): Note | ErrorResponse;

@route("/api/v1/notes/search")
@get
@summary("Search notes")
@tag("Notes")
op searchNotes(
  @query query: string
): SearchNotesResponse | ErrorResponse;

@route("/api/v1/notes/{noteId}/tags")
@post
@summary("Add tags to a note")
@tag("Notes")
op addTagsToNote(
  @path noteId: string,
  @body requestBody: AddTagsRequest
): Note | ErrorResponse;

@route("/api/v1/notes/{noteId}/images")
@post
@summary("Upload an image to a note")
@tag("Notes")
op uploadNoteImage(
  @path noteId: string,
  @body requestBody: UploadImageRequest
): Note | ErrorResponse;

@route("/api/v1/note-templates")
@get
@summary("Get note templates")
@tag("Notes")
op getNoteTemplates(): string[] | ErrorResponse;

@route("/api/v1/notes/{noteId}/draft")
@post
@summary("Save a note as draft")
@tag("Notes")
op saveNoteDraft(
  @path noteId: string
): Note | ErrorResponse;

model CalendarEvent {
  title: string;
  date: utcDateTime;
}

model NoteSummary {
  count: int32;
  notes: Note[];
}

@route("/api/v1/analytics/daily")
@get
@summary("Get daily note summary")
@tag("Analytics")
op getDailySummary(): NoteSummary | ErrorResponse;

@route("/api/v1/analytics/weekly")
@get
@summary("Get weekly note summary")
@tag("Analytics")
op getWeeklySummary(): NoteSummary | ErrorResponse;

@route("/api/v1/analytics/monthly")
@get
@summary("Get monthly note summary")
@tag("Analytics")
op getMonthlySummary(): NoteSummary | ErrorResponse;

@route("/api/v1/analytics/yearly")
@get
@summary("Get yearly note summary")
@tag("Analytics")
op getYearlySummary(): NoteSummary | ErrorResponse;

@route("/api/v1/analytics/quarterly")
@get
@summary("Get quarterly note summary")
@tag("Analytics")
op getQuarterlySummary(): NoteSummary | ErrorResponse;

model TagAnalytics {
  tag: string;
  count: int32;
}

model GetTagAnalyticsResponse {
  tags: TagAnalytics[];
}

@route("/api/v1/analytics/tags")
@get
@summary("Get tag usage analytics")
@tag("Analytics")
op getTagAnalytics(): GetTagAnalyticsResponse | ErrorResponse;

model IntegrationService {
  serviceName: string;
  isConnected: boolean;
  data?: string;
}

@route("/api/v1/integrations/{service}/connect")
@post
@summary("Connect to an external service")
@tag("Integrations")
op connectService(
  @path service: string
): void | ErrorResponse;

@route("/api/v1/integrations/{service}/disconnect")
@delete
@summary("Disconnect from an external service")
@tag("Integrations")
op disconnectService(
  @path service: string
): void | ErrorResponse;

@route("/api/v1/integrations/{service}/data")
@get
@summary("Get data from a connected service")
@tag("Integrations")
op getServiceData(
  @path service: string
): IntegrationService | ErrorResponse;

model ShareNoteRequest {
  shareWith: string;
}

@route("/api/v1/notes/{noteId}/share")
@post
@summary("Share a note with another user")
@tag("Notes")
op shareNote(
  @path noteId: string,
  @body requestBody: ShareNoteRequest
): void | ErrorResponse;

@route("/api/v1/notes/{noteId}/share")
@delete
@summary("Unshare a previously shared note")
@tag("Notes")
op unshareNote(
  @path noteId: string
): void | ErrorResponse;

model GetAllTagsResponse {
  tags: string[];
}

@route("/api/v1/tags")
@get
@summary("Get all tags")
@tag("Tags")
op getAllTags(): GetAllTagsResponse | ErrorResponse;

@route("/api/v1/health")
@get
@summary("Check API health")
@tag("Health")
op healthCheck(): string | ErrorResponse;
