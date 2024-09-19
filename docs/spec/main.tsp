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
@server("https://chronotes.yashikota.com/mock", "mock")
@server("https://chronotes.yashikota.com", "prd")
namespace chronotes;

enum Versions {
  v1,
}

@example(#{
    id: "1",
    username: "JohnDoe",
    email: "john.doe@example.com",
    createdAt: utcDateTime.fromISO("2021-01-01T00:00:00Z"),
})
model User {
  id: string;
  username: string;
  email: string;
  createdAt: utcDateTime;
}

@example(#{
    username: "JohnDoe",
    email: "john.doe@example.com",
    password: "password",
})
model RegisterUserRequest {
  username: string;
  email: string;
  password: string;
}

@example(#{
  email: "john.doe@example.com",
  password: "password123"
})
model LoginRequest {
  email: string;
  password: string;
}

@example(#{
  user: #{
    id: "1",
    username: "JohnDoe",
    email: "john.doe@example.com",
    createdAt: utcDateTime.fromISO("2021-01-01T00:00:00Z"),
  },
  token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG4uZG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
})
model RegisterUserResponse {
  user: User;
  token: string;
}

@example(#{
  message: "Error message"
})
@error
model ErrorResponse {
  message: string;
}

@route("/api/v1/users/register")
@post
@summary("Register a new user")
@tag("Users")
@useAuth(NoAuth)
@opExample(#{
    parameters: #{
        requestBody: #{
        username: "JohnDoe",
        email: "john.doe@example.com",
        password: "password",
    }},
    returnType: #{
        user: #{
            id: "1",
            username: "JohnDoe",
            email: "john.doe@example.com",
            createdAt: utcDateTime.fromISO("2021-01-01T00:00:00Z"),
        },
        token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG4uZG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
    }
})
op registerUser(
  @body requestBody: RegisterUserRequest
): RegisterUserResponse | ErrorResponse;

@route("/api/v1/users/{userId}")
@delete
@summary("Delete a user")
@tag("Users")
@useAuth(BearerAuth)
@opExample(#{
  parameters: #{
    userId: "user123"
  },
})
op deleteUser(
  @path userId: string
): void | ErrorResponse;

@route("/api/v1/users/login")
@post
@summary("Login user")
@tag("Users")
@useAuth(BearerAuth)
@opExample(#{
  parameters: #{
    requestBody: #{
      email: "john.doe@example.com",
      password: "password123"
    }
  },
  returnType: #{
    user: #{
      id: "1",
      username: "JohnDoe",
      email: "john.doe@example.com",
      createdAt: utcDateTime.fromISO("2021-01-01T00:00:00Z"),
    },
    token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG4uZG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
  }
})
op login(
  @body requestBody: LoginRequest
): RegisterUserResponse | ErrorResponse;

@route("/api/v1/users/logout")
@post
@summary("Logout user")
@tag("Users")
@useAuth(BearerAuth)
@opExample(#{
  parameters: #{},
})
op logout(): void | ErrorResponse;

@example(#{
  id: "note123",
  title: "My First Note",
  content: "This is the content of my first note.",
  tags: #["personal", "ideas"],
  images: #[
    #{
      id: "img1",
      url: "https://example.com/image1.jpg"
    }
  ],
  isDraft: false,
  createdAt: utcDateTime.fromISO("2023-01-01T12:00:00Z"),
  updatedAt: utcDateTime.fromISO("2023-01-01T12:30:00Z")
})
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

@example(#{
  title: "My New Note",
  content: "This is the content of my new note.",
  tags: #["work", "project"],
  images: #[
    #{
      id: "img2",
      url: "https://example.com/image2.jpg"
    }
  ],
  isDraft: true
})
model CreateNoteRequest {
  title: string;
  content: string;
  tags?: string[];
  images?: Image[];
  isDraft?: boolean;
}

@example(#{
  title: "Updated Note Title",
  content: "This is the updated content of my note.",
  tags: #["personal", "updated"],
  images: #[
    #{
      id: "img3",
      url: "https://example.com/image3.jpg"
    }
  ],
  isDraft: false
})
model UpdateNoteRequest {
  title?: string;
  content?: string;
  tags?: string[];
  images?: Image[];
  isDraft?: boolean;
}

@example(#{
  notes: #[
    #{
      id: "note1",
      title: "First Note",
      content: "Content of first note",
      tags: #["personal"],
      images: #[
        #{
        id: "img2",
        url: "https://example.com/image2.jpg"
        }
      ],
      isDraft: false,
      createdAt: utcDateTime.fromISO("2023-01-01T10:00:00Z"),
      updatedAt: utcDateTime.fromISO("2023-01-01T10:00:00Z")
    },
    #{
      id: "note2",
      title: "Second Note",
      content: "Content of second note",
      tags: #["work"],
      images: #[
        #{
        id: "img2",
        url: "https://example.com/image2.jpg"
        }
      ],
      isDraft: true,
      createdAt: utcDateTime.fromISO("2023-01-02T14:30:00Z"),
      updatedAt: utcDateTime.fromISO("2023-01-02T14:30:00Z")
    }
  ],
  totalCount: 2
})
model GetNotesResponse {
  notes: Note[];
  totalCount: int32;
}

@example(#{
  notes: #[
    #{
      id: "note3",
      title: "Matching Note",
      content: "This note matches the search query",
      tags: #["search", "result"],
      images: #[
        #{
        id: "img2",
        url: "https://example.com/image2.jpg"
        }
      ],
      isDraft: false,
      createdAt: utcDateTime.fromISO("2023-01-03T09:15:00Z"),
      updatedAt: utcDateTime.fromISO("2023-01-03T09:15:00Z")
    }
  ]
})
model SearchNotesResponse {
  notes: Note[];
}

@example(#{
  tags: #["important", "urgent", "project"]
})
model AddTagsRequest {
  tags: string[];
}

@example(#{
  image: #{
    id: "img4",
    url: "https://example.com/image4.jpg"
  }
})
model UploadImageRequest {
  image: Image;
}

@example(#{
  id: "img5",
  url: "https://example.com/image5.jpg"
})
model Image {
  id: string;
  url: string;
}

@route("/api/v1/notes")
@post
@summary("Create a new note")
@tag("Notes")
@useAuth(BearerAuth)
@opExample(#{
  parameters: #{
    requestBody: #{
      title: "My New Note",
      content: "This is the content of my new note.",
      tags: #["work", "project"],
      images: #[
        #{
          id: "img2",
          url: "https://example.com/image2.jpg"
        }
      ],
      isDraft: true
    }
  },
  returnType: #{
    id: "note123",
    title: "My New Note",
    content: "This is the content of my new note.",
    tags: #["work", "project"],
    images: #[
      #{
        id: "img2",
        url: "https://example.com/image2.jpg"
      }
    ],
    isDraft: true,
    createdAt: utcDateTime.fromISO("2023-01-01T12:00:00Z"),
    updatedAt: utcDateTime.fromISO("2023-01-01T12:00:00Z")
  }
})
op createNote(
  @body requestBody: CreateNoteRequest
): Note | ErrorResponse;

@route("/api/v1/notes/{noteId}")
@put
@summary("Update an existing note")
@tag("Notes")
@useAuth(BearerAuth)
@opExample(#{
  parameters: #{
    noteId: "note123",
    requestBody: #{
      title: "Updated Note Title",
      content: "This is the updated content of my note.",
      tags: #["personal", "updated"],
      images: #[
        #{
          id: "img3",
          url: "https://example.com/image3.jpg"
        }
      ],
      isDraft: false
    }
  },
  returnType: #{
    id: "note123",
    title: "Updated Note Title",
    content: "This is the updated content of my note.",
    tags: #["personal", "updated"],
    images: #[
      #{
        id: "img3",
        url: "https://example.com/image3.jpg"
      }
    ],
    isDraft: false,
    createdAt: utcDateTime.fromISO("2023-01-01T12:00:00Z"),
    updatedAt: utcDateTime.fromISO("2023-01-01T13:00:00Z")
  }
})
op updateNote(
  @path noteId: string,
  @body requestBody: UpdateNoteRequest
): Note | ErrorResponse;

@route("/api/v1/notes/{noteId}")
@delete
@summary("Delete a note")
@tag("Notes")
@useAuth(BearerAuth)
@opExample(#{
  parameters: #{
    noteId: "note123"
  },
})
op deleteNote(
  @path noteId: string
): void | ErrorResponse;

@route("/api/v1/notes")
@get
@summary("Get a list of notes")
@tag("Notes")
@useAuth(BearerAuth)
@opExample(#{
  parameters: #{
    page: 1,
    pageSize: 10
  },
  returnType: #{
    notes: #[
      #{
        id: "note1",
        title: "First Note",
        content: "Content of first note",
        tags: #["personal"],
        images: #[
          #{
            id: "img1",
            url: "https://example.com/image1.jpg"
          }
        ],
        isDraft: false,
        createdAt: utcDateTime.fromISO("2023-01-01T10:00:00Z"),
        updatedAt: utcDateTime.fromISO("2023-01-01T10:00:00Z")
      },
      #{
        id: "note2",
        title: "Second Note",
        content: "Content of second note",
        tags: #["work"],
        images: #[],
        isDraft: true,
        createdAt: utcDateTime.fromISO("2023-01-02T14:30:00Z"),
        updatedAt: utcDateTime.fromISO("2023-01-02T14:30:00Z")
      }
    ],
    totalCount: 2
  }
})
op getNotes(
  @query page?: int32,
  @query pageSize?: int32
): GetNotesResponse | ErrorResponse;

@route("/api/v1/notes/{noteId}")
@get
@summary("Get a specific note")
@tag("Notes")
@useAuth(BearerAuth)
@opExample(#{
  parameters: #{
    noteId: "note123"
  },
  returnType: #{
    id: "note123",
    title: "My First Note",
    content: "This is the content of my first note.",
    tags: #["personal", "ideas"],
    images: #[
      #{
        id: "img1",
        url: "https://example.com/image1.jpg"
      }
    ],
    isDraft: false,
    createdAt: utcDateTime.fromISO("2023-01-01T12:00:00Z"),
    updatedAt: utcDateTime.fromISO("2023-01-01T12:30:00Z")
  }
})
op getNote(
  @path noteId: string
): Note | ErrorResponse;

@route("/api/v1/notes/search")
@get
@summary("Search notes")
@tag("Notes")
@useAuth(BearerAuth)
@opExample(#{
  parameters: #{
    query: "important"
  },
  returnType: #{
    notes: #[
      #{
        id: "note3",
        title: "Matching Note",
        content: "This note matches the search query",
        tags: #["search", "result"],
        images: #[],
        isDraft: false,
        createdAt: utcDateTime.fromISO("2023-01-03T09:15:00Z"),
        updatedAt: utcDateTime.fromISO("2023-01-03T09:15:00Z")
      }
    ]
  }
})
op searchNotes(
  @query query: string
): SearchNotesResponse | ErrorResponse;

@route("/api/v1/notes/{noteId}/tags")
@post
@summary("Add tags to a note")
@tag("Notes")
@useAuth(BearerAuth)
@opExample(#{
  parameters: #{
    noteId: "note123",
    requestBody: #{
      tags: #["important", "urgent", "project"]
    }
  },
  returnType: #{
    id: "note123",
    title: "My First Note",
    content: "This is the content of my first note.",
    tags: #["personal", "ideas", "important", "urgent", "project"],
    images: #[
      #{
        id: "img1",
        url: "https://example.com/image1.jpg"
      }
    ],
    isDraft: false,
    createdAt: utcDateTime.fromISO("2023-01-01T12:00:00Z"),
    updatedAt: utcDateTime.fromISO("2023-01-01T14:00:00Z")
  }
})
op addTagsToNote(
  @path noteId: string,
  @body requestBody: AddTagsRequest
): Note | ErrorResponse;

@route("/api/v1/notes/{noteId}/images")
@post
@summary("Upload an image to a note")
@tag("Notes")
@useAuth(BearerAuth)
@opExample(#{
  parameters: #{
    noteId: "note123",
    requestBody: #{
      image: #{
        id: "img4",
        url: "https://example.com/image4.jpg"
      }
    }
  },
  returnType: #{
    id: "note123",
    title: "My First Note",
    content: "This is the content of my first note.",
    tags: #["personal", "ideas"],
    images: #[
      #{
        id: "img1",
        url: "https://example.com/image1.jpg"
      },
      #{
        id: "img4",
        url: "https://example.com/image4.jpg"
      }
    ],
    isDraft: false,
    createdAt: utcDateTime.fromISO("2023-01-01T12:00:00Z"),
    updatedAt: utcDateTime.fromISO("2023-01-01T15:00:00Z")
  }
})
op uploadNoteImage(
  @path noteId: string,
  @body requestBody: UploadImageRequest
): Note | ErrorResponse;

@route("/api/v1/note-templates")
@get
@summary("Get note templates")
@tag("Notes")
@useAuth(BearerAuth)
@opExample(#{
  parameters: #{},
  returnType: #["Meeting Notes", "Daily Journal", "Project Plan"]
})
op getNoteTemplates(): string[] | ErrorResponse;

@route("/api/v1/notes/{noteId}/draft")
@post
@summary("Save a note as draft")
@tag("Notes")
@useAuth(BearerAuth)
@opExample(#{
  parameters: #{
    noteId: "note123"
  },
  returnType: #{
    id: "note123",
    title: "My First Note",
    content: "This is the content of my first note.",
    tags: #["personal", "ideas"],
    images: #[
      #{
        id: "img1",
        url: "https://example.com/image1.jpg"
      }
    ],
    isDraft: true,
    createdAt: utcDateTime.fromISO("2023-01-01T12:00:00Z"),
    updatedAt: utcDateTime.fromISO("2023-01-01T16:00:00Z")
  }
})
op saveNoteDraft(
  @path noteId: string
): Note | ErrorResponse;

@example(#{
  count: 5,
  notes: #[
    #{
      id: "note1",
      title: "Summary Note 1",
      content: "Content of summary note 1",
      tags: #["summary"],
      images: #[
        #{
        id: "img2",
        url: "https://example.com/image2.jpg"
        }
      ],
      isDraft: false,
      createdAt: utcDateTime.fromISO("2023-01-01T08:00:00Z"),
      updatedAt: utcDateTime.fromISO("2023-01-01T08:00:00Z")
    },
    #{
      id: "note2",
      title: "Summary Note 2",
      content: "Content of summary note 2",
      tags: #["summary"],
      images: #[
        #{
        id: "img2",
        url: "https://example.com/image2.jpg"
        }
      ],
      isDraft: false,
      createdAt: utcDateTime.fromISO("2023-01-02T10:30:00Z"),
      updatedAt: utcDateTime.fromISO("2023-01-02T10:30:00Z")
    }
  ]
})
model NoteSummary {
  count: int32;
  notes: Note[];
}

@route("/api/v1/analytics/daily")
@get
@summary("Get daily note summary")
@tag("Analytics")
@useAuth(BearerAuth)
@opExample(#{
  parameters: #{},
  returnType: #{
    count: 3,
    notes: #[
      #{
        id: "note1",
        title: "Daily Summary Note 1",
        content: "Content of daily summary note 1",
        tags: #["daily", "summary"],
        images: #[],
        isDraft: false,
        createdAt: utcDateTime.fromISO("2023-01-01T08:00:00Z"),
        updatedAt: utcDateTime.fromISO("2023-01-01T08:00:00Z")
      },
      #{
        id: "note2",
        title: "Daily Summary Note 2",
        content: "Content of daily summary note 2",
        tags: #["daily", "summary"],
        images: #[],
        isDraft: false,
        createdAt: utcDateTime.fromISO("2023-01-01T10:30:00Z"),
        updatedAt: utcDateTime.fromISO("2023-01-01T10:30:00Z")
      }
    ]
  }
})
op getDailySummary(): NoteSummary | ErrorResponse;

@route("/api/v1/analytics/weekly")
@get
@summary("Get weekly note summary")
@tag("Analytics")
@useAuth(BearerAuth)
@opExample(#{
  parameters: #{},
  returnType: #{
    count: 10,
    notes: #[
      #{
        id: "note1",
        title: "Weekly Summary Note 1",
        content: "Content of weekly summary note 1",
        tags: #["weekly", "summary"],
        images: #[],
        isDraft: false,
        createdAt: utcDateTime.fromISO("2023-01-01T08:00:00Z"),
        updatedAt: utcDateTime.fromISO("2023-01-01T08:00:00Z")
      },
      #{
        id: "note2",
        title: "Weekly Summary Note 2",
        content: "Content of weekly summary note 2",
        tags: #["weekly", "summary"],
        images: #[],
        isDraft: false,
        createdAt: utcDateTime.fromISO("2023-01-02T10:30:00Z"),
        updatedAt: utcDateTime.fromISO("2023-01-02T10:30:00Z")
      }
    ]
  }
})
op getWeeklySummary(): NoteSummary | ErrorResponse;

@route("/api/v1/analytics/monthly")
@get
@summary("Get monthly note summary")
@tag("Analytics")
@useAuth(BearerAuth)
@opExample(#{
  parameters: #{},
  returnType: #{
    count: 25,
    notes: #[
      #{
        id: "note1",
        title: "Monthly Summary Note 1",
        content: "Content of monthly summary note 1",
        tags: #["monthly", "summary"],
        images: #[],
        isDraft: false,
        createdAt: utcDateTime.fromISO("2023-01-01T08:00:00Z"),
        updatedAt: utcDateTime.fromISO("2023-01-01T08:00:00Z")
      },
      #{
        id: "note2",
        title: "Monthly Summary Note 2",
        content: "Content of monthly summary note 2",
        tags: #["monthly", "summary"],
        images: #[],
        isDraft: false,
        createdAt: utcDateTime.fromISO("2023-01-15T10:30:00Z"),
        updatedAt: utcDateTime.fromISO("2023-01-15T10:30:00Z")
      }
    ]
  }
})
op getMonthlySummary(): NoteSummary | ErrorResponse;

@route("/api/v1/analytics/yearly")
@get
@summary("Get yearly note summary")
@tag("Analytics")
@useAuth(BearerAuth)
@opExample(#{
  parameters: #{},
  returnType: #{
    count: 150,
    notes: #[
      #{
        id: "note1",
        title: "Yearly Summary Note 1",
        content: "Content of yearly summary note 1",
        tags: #["yearly", "summary"],
        images: #[],
        isDraft: false,
        createdAt: utcDateTime.fromISO("2023-01-01T08:00:00Z"),
        updatedAt: utcDateTime.fromISO("2023-01-01T08:00:00Z")
      },
      #{
        id: "note2",
        title: "Yearly Summary Note 2",
        content: "Content of yearly summary note 2",
        tags: #["yearly", "summary"],
        images: #[],
        isDraft: false,
        createdAt: utcDateTime.fromISO("2023-06-15T10:30:00Z"),
        updatedAt: utcDateTime.fromISO("2023-06-15T10:30:00Z")
      }
    ]
  }
})
op getYearlySummary(): NoteSummary | ErrorResponse;

@route("/api/v1/analytics/quarterly")
@get
@summary("Get quarterly note summary")
@tag("Analytics")
@useAuth(BearerAuth)
@opExample(#{
  parameters: #{},
  returnType: #{
    count: 75,
    notes: #[
      #{
        id: "note1",
        title: "Quarterly Summary Note 1",
        content: "Content of quarterly summary note 1",
        tags: #["quarterly", "summary"],
        images: #[],
        isDraft: false,
        createdAt: utcDateTime.fromISO("2023-01-01T08:00:00Z"),
        updatedAt: utcDateTime.fromISO("2023-01-01T08:00:00Z")
      },
      #{
        id: "note2",
        title: "Quarterly Summary Note 2",
        content: "Content of quarterly summary note 2",
        tags: #["quarterly", "summary"],
        images: #[],
        isDraft: false,
        createdAt: utcDateTime.fromISO("2023-03-15T10:30:00Z"),
        updatedAt: utcDateTime.fromISO("2023-03-15T10:30:00Z")
      }
    ]
  }
})
op getQuarterlySummary(): NoteSummary | ErrorResponse;

@example(#{
  tag: "project",
  count: 7
})
model TagAnalytics {
  tag: string;
  count: int32;
}

@example(#{
  tags: #[
    #{
      tag: "work",
      count: 10
    },
    #{
      tag: "personal",
      count: 5
    },
    #{
      tag: "ideas",
      count: 3
    }
  ]
})
model GetTagAnalyticsResponse {
  tags: TagAnalytics[];
}

@route("/api/v1/analytics/tags")
@get
@summary("Get tag usage analytics")
@tag("Analytics")
@useAuth(BearerAuth)
@opExample(#{
  parameters: #{},
  returnType: #{
    tags: #[
      #{
        tag: "work",
        count: 50
      },
      #{
        tag: "personal",
        count: 30
      },
      #{
        tag: "ideas",
        count: 20
      },
      #{
        tag: "project",
        count: 40
      }
    ]
  }
})
op getTagAnalytics(): GetTagAnalyticsResponse | ErrorResponse;

@example(#{
  serviceName: "Google Drive",
  isConnected: true,
  data: "Last synced: 2023-01-15T14:30:00Z"
})
model IntegrationService {
  serviceName: string;
  isConnected: boolean;
  data?: string;
}

@route("/api/v1/integrations/{service}/connect")
@post
@summary("Connect to an external service")
@tag("Integrations")
@useAuth(BearerAuth)
@opExample(#{
  parameters: #{
    service: "google-drive"
  },
})
op connectService(
  @path service: string
): void | ErrorResponse;

@route("/api/v1/integrations/{service}/disconnect")
@delete
@summary("Disconnect from an external service")
@tag("Integrations")
@useAuth(BearerAuth)
@opExample(#{
  parameters: #{
    service: "google-drive"
  },
})
op disconnectService(
  @path service: string
): void | ErrorResponse;

@route("/api/v1/integrations/{service}/data")
@get
@summary("Get data from a connected service")
@tag("Integrations")
@useAuth(BearerAuth)
@opExample(#{
  parameters: #{
    service: "google-drive"
  },
  returnType: #{
    serviceName: "Google Drive",
    isConnected: true,
    data: "Last synced: 2023-01-15T14:30:00Z"
  }
})
op getServiceData(
  @path service: string
): IntegrationService | ErrorResponse;

@example(#{
  shareWith: "jane.doe@example.com"
})
model ShareNoteRequest {
  shareWith: string;
}

@route("/api/v1/notes/{noteId}/share")
@post
@summary("Share a note with another user")
@tag("Notes")
@useAuth(BearerAuth)
@opExample(#{
  parameters: #{
    noteId: "note123",
    requestBody: #{
      shareWith: "jane.doe@example.com"
    }
  },
})
op shareNote(
  @path noteId: string,
  @body requestBody: ShareNoteRequest
): void | ErrorResponse;

@route("/api/v1/notes/{noteId}/share")
@delete
@summary("Unshare a previously shared note")
@tag("Notes")
@useAuth(BearerAuth)
@opExample(#{
  parameters: #{
    noteId: "note123"
  },
})
op unshareNote(
  @path noteId: string
): void | ErrorResponse;

@example(#{
  tags: #["work", "personal", "ideas", "project", "urgent"]
})
model GetAllTagsResponse {
  tags: string[];
}

@route("/api/v1/tags")
@get
@summary("Get all tags")
@tag("Tags")
@useAuth(BearerAuth)
@opExample(#{
  parameters: #{},
  returnType: #{
    tags: #["work", "personal", "ideas", "project", "urgent", "summary", "daily", "weekly", "monthly", "yearly", "quarterly"]
  }
})
op getAllTags(): GetAllTagsResponse | ErrorResponse;

@route("/api/v1/health")
@get
@summary("Check API health")
@tag("Health")
@useAuth(NoAuth)
@opExample(#{
  parameters: #{},
  returnType: "pong"
})
op healthCheck(): string | ErrorResponse;
