using TypeSpec.Http;

namespace Chronotes {
  @route("/api/v1/admin")
  @tag("Admin")
  @useAuth(BearerAuth)
  interface Admin {
    @route("/notes")
    @post
    @summary("Add notes")
    addNotes(@body requestBody: Note[]): SuccessResponse | ErrorResponse;

    @route("/notes")
    @delete
    @summary("Delete notes")
    deleteNotes(@query user_id: UserID): SuccessResponse | ErrorResponse;
  }
}
