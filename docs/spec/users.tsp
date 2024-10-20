using TypeSpec.Http;

namespace Chronotes {
  model GetUpdateUserResponse {
    user_id: User.user_id;
    user_name: User.user_name;
    email: User.email;
    createdAt: User.createdAt;
    updatedAt: User.updatedAt;
    accounts: Accounts;
  }

  model UpdateUserRequest {
    user_name?: User.user_name;
    email?: User.email;
    password?: User.password;
    accounts?: Accounts;
  }

  @route("/api/v1/users/me")
  @tag("Users")
  @useAuth(BearerAuth)
  interface Users {
    @get
    @summary("Get user information")
    get(): GetUpdateUserResponse | ErrorResponse;

    /**
     * 更新したい要素のみ指定すること \
     * 例えば、GitHubアカウントのみ更新する場合
     * ```json
     * {
     *   "accounts": {
     *   "github_user_id": "yashikota"
     * }
     * ```
     * このリクエストを送信する
     */
    @put
    @summary("Update user information")
    update(
      @body requestBody: UpdateUserRequest,
    ): GetUpdateUserResponse | ErrorResponse;

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
      @body requestBody: {
        password: User.password;
      },
    ): SuccessResponse | ErrorResponse;
  }
}
