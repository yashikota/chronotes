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
     * - `tags` \
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

  @route("/api/v1/notes/share")
  @tag("Notes")
  @useAuth(BearerAuth)
  interface ShareNotes {
    @post
    @summary("Share notes")
    shareNotes(@query note_id: Note.note_id): SuccessResponse | ErrorResponse;

    @get
    @summary("Get shared notes")
    getSharedNotes(@query share_url: Note.share_url): Note | ErrorResponse;

    @delete
    @summary("Unshare notes")
    unshareNotes(@query note_id: Note.note_id): SuccessResponse | ErrorResponse;
  }
}
