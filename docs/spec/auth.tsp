using TypeSpec.Http;

namespace Chronotes {
  model RegisterRequest {
    user_id: User.user_id;
    user_name: User.user_name;
    email: User.email;
    password: User.password;
  }

  model LoginRequest {
    user_id?: User.user_id;
    email?: User.email;
    password: User.password;
  }

  // Route
  @route("/api/v1/auth")
  @tag("Auth")
  @useAuth(NoAuth)
  interface Auth {
    @route("/register")
    @post
    @summary("Register a new user")
    register(@body requestBody: RegisterRequest): {
      token: User.token;
    } | ErrorResponse;

    /**
     * `user_id` または `email` のどちらかでログイン可能 \
     * `user_id` でログインする場合は、
     * ```json
     * {
     *   "user_id": "yashikota",
     *   "password": "password"
     * }
     * ```
     *
     * `email` でログインする場合は、
     * ```json
     * {
     *   "email": "test@example.com",
     *   "password": "password"
     * }
     * ```
     */
    @route("/login")
    @post
    @summary("Login user")
    login(@body requestBody: LoginRequest): {
      token: User.token;
    } | ErrorResponse;

    @route("/logout")
    @post
    @summary("Logout user")
    @useAuth(BearerAuth)
    logout(): SuccessResponse | ErrorResponse;
  }
}
