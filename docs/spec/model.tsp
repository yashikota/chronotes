namespace Chronotes {
  // User
  model User {
    @doc("任意に変更可能なID。英数字とアンダースコアのみ。4文字以上20文字以下")
    @pattern("^[a-zA-Z0-9_]+$")
    @minLength(4)
    @maxLength(20)
    @example("test123")
    user_id: string;

    @doc("ユーザー名。1文字以上20文字以下")
    @minLength(1)
    @maxLength(20)
    @example("hoge太郎")
    user_name: string;

    @doc("メールアドレス")
    @format("email")
    @pattern("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$")
    @example("test@example.com")
    email: string;

    @doc("パスワード。8文字以上20文字以下")
    @minLength(8)
    @maxLength(20)
    @example("password")
    password: string;

    @doc("JWTトークン")
    @format("jwt")
    @pattern("^[a-zA-Z0-9-_=]+\\.[a-zA-Z0-9-_=]+\\.[a-zA-Z0-9-_.+/=]+$")
    @example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG4uZG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")
    token: string;

    @doc("作成日時")
    @example(utcDateTime.fromISO("2024-10-01T19:30:06Z"))
    createdAt: utcDateTime;

    @doc("更新日時")
    @example(utcDateTime.fromISO("2024-10-02T03:34:56Z"))
    updatedAt: utcDateTime;

    @doc("外部アカウント情報")
    accounts?: Accounts;
  }

  model Accounts {
    @doc("GitHubのユーザーID")
    @pattern("^[a-zA-Z0-9-]+$")
    @example("yashikota")
    github_user_id?: string;

    @doc("DiscordのチャンネルID")
    @pattern("^[0-9]+$")
    @example("9419543265265028")
    discord_channel_id?: string;

    @doc("SlackのチャンネルID")
    @pattern("^[a-zA-Z0-9]+$")
    @example("HG824ORHQ250")
    slack_channel_id?: string;

    @doc("QiitaのユーザーID")
    @pattern("^[a-zA-Z0-9_]+$")
    @example("yashikota")
    qiita_user_id?: string;

    @doc("Zennのユーザー名")
    @pattern("^[a-zA-Z0-9_]+$")
    @example("yashikota")
    zenn_username?: string;

    @doc("Wakatimeのトークン")
    @pattern("^[a-z0-9-]+$")
    @example("waka_763hgfds1-2ushy-2yfds-lkaxvczv-63hdfgshvds")
    wakatime_token?: string;
  }

  // Note
  model Note {
    user_id: User.user_id;

    @doc("内部で付与されるID")
    @pattern("^[a-zA-Z0-9]+$")
    @example("FHNWFLAQN")
    note_id: string;

    @doc("タイトル")
    @minLength(1)
    @maxLength(40)
    @example("GoとPythonの比較")
    title: string;

    @doc("内容")
    @minLength(1)
    @maxLength(4000)
    @example("<h1>GoとPythonの比較</h1><p>GoとPythonの比較を行いました。</p>")
    content: string;

    @doc("タグ。`,`区切りで複数指定可能")
    @pattern("^[a-z0-9, ]+$", "小文字英数字のみ")
    @example("go,python")
    tags: string;

    @doc("文字数")
    length: integer;

    @doc("共有URL")
    @format("share_url")
    @example("01J9ESHFPMJDX5FH4JKQKE6TJM")
    share_url?: string;

    createdAt: User.createdAt;
    updatedAt: User.updatedAt;
  }
}
