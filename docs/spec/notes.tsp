using TypeSpec.Http;

namespace Chronotes {
  const noteExample = #{
    user_id: UserIDExample,
    note_id: "FHNWFLAQN",
    title: "My First Note",
    content: "This is the content of my first note.",
    tags: "work, project",
    createdAt: utcDateTime.fromISO("2023-01-01T12:00:00Z"),
    updatedAt: utcDateTime.fromISO("2023-01-01T12:30:00Z"),
  };
  @example(noteExample)
  model Note {
    user_id: UserID;
    note_id: string;
    title: string;
    content: string;
    tags: string;
    createdAt: utcDateTime;
    updatedAt: utcDateTime;
  }

  @example(#{ notes: #[noteExample, noteExample] })
  model SearchNotesResponse {
    notes: Note[];
  }

  @route("/api/v1/notes")
  @tag("Notes")
  @useAuth(BearerAuth)
  interface Notes {
    @get
    @summary("Get a list of notes")
    @doc("fieldsは `note_id`, `user_id`, `title`, `content`, `length`, `tags`が指定可能。複数指定する場合は空白なしのコンマで指定 (例) `title,tags`")
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
    searchNotes(@query query: string): SearchNotesResponse | ErrorResponse;
  }
}
