using TypeSpec.Http;

namespace Chronotes {
  @route("/api/v1/notes")
  @tag("Notes")
  @useAuth(BearerAuth)
  interface Notes {
    /**
     * fieldsは
     * - `note_id`
     * - `user_id`
     * - `title`
     * - `content`
     * - `length`
     * - `tags`
     * - `created_at`
     * - `updated_at` \
     * が指定可能。\
     * 複数指定する場合は空白なしのコンマで指定 \
     * (例) `title,tags`
     */
    @get
    @summary("Get a list of notes")
    getNotes(
      @query from: utcDateTime,
      @query to: utcDateTime,
      @query fields: string,
    ): Note[] | ErrorResponse;

    @put
    @summary("Update a notes")
    updateNote(@body requestBody: Note[]): Note | ErrorResponse;

    @delete
    @summary("Delete a notes")
    deleteNote(@body requestBody: Note[]): SuccessResponse | ErrorResponse;

    @route("/search")
    @post
    @summary("Search notes")
    searchNotes(@query query: string): Note[] | ErrorResponse;
  }
}
