using TypeSpec.Http;

namespace Chronotes {
  // Models
  const UserIDExample = #{ user_id: "test123" };
  model UserID {
    @doc("任意に変更可能なID。英数字とアンダースコアのみ。4文字以上20文字以下")
    @pattern("^[a-zA-Z0-9_]+$")
    @minLength(4)
    @maxLength(20)
    user_id: string;
  }

  const UserNameExample = #{ user_name: "hoge太郎" };
  model UserName {
    @doc("ユーザー名。1文字以上20文字以下")
    @minLength(1)
    @maxLength(20)
    user_name: string;
  }

  const MailAddressExample = #{ email: "test@example.com" };
  model MailAddress {
    @doc("メールアドレス")
    @format("email")
    @pattern("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$")
    email: string;
  }

  const PasswordExample = #{ password: "password" };
  model Password {
    @doc("パスワード。8文字以上20文字以下")
    @minLength(8)
    @maxLength(20)
    password: string;
  }

  const tokenExample = #{
    token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG4uZG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
  };
  model Token {
    @doc("JWTトークン")
    @format("jwt")
    @pattern("^[a-zA-Z0-9-_=]+\\.[a-zA-Z0-9-_=]+\\.[a-zA-Z0-9-_.+/=]+$")
    token: string;
  }

  // Register
  const RegisterUserRequestExample = #{
    user_id: UserIDExample,
    user_name: UserNameExample,
    email: MailAddressExample,
    password: PasswordExample,
  };
  @example(RegisterUserRequestExample)
  model RegisterRequest {
    user_id: UserID;
    user_name: UserName;
    email: MailAddress;
    password: Password;
  }

  const RegisterUserResponseExample = #{ token: tokenExample };
  model RegisterLoginResponse {
    token: Token;
  }

  // Login
  const LoginRequestExample = #{
    user_id: UserIDExample,
    email: MailAddressExample,
    password: PasswordExample,
  };
  @example(LoginRequestExample)
  @doc("`user_id` または `email` のどちらかでログイン可能")
  model LoginRequest {
    user_id?: UserID;
    email?: MailAddress;
    password: Password;
  }

  // Route
  @route("/api/v1/auth")
  @tag("Auth")
  @useAuth(NoAuth)
  interface Auth {
    @route("/register")
    @post
    @summary("Register a new user")
    register(
      @body requestBody: RegisterRequest,
    ): RegisterLoginResponse | ErrorResponse;

    @route("/login")
    @post
    @summary("Login user")
    login(
      @body requestBody: LoginRequest,
    ): RegisterLoginResponse | ErrorResponse;

    @route("/logout")
    @post
    @summary("Logout user")
    @useAuth(BearerAuth)
    logout(): SuccessResponse | ErrorResponse;
  }
}
