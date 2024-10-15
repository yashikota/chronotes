using TypeSpec.Http;

namespace Chronotes {
  const userExample = #{
    user_id: UserIDExample,
    user_name: UserNameExample,
    email: MailAddressExample,
    createdAt: utcDateTime.fromISO("2024-10-01T19:30:06Z"),
    updatedAt: utcDateTime.fromISO("2024-10-02T03:34:56Z"),
    accounts: AccountsExample,
  };
  @example(userExample)
  model User {
    user_id: UserID;
    user_name: UserName;
    email: MailAddress;
    createdAt: utcDateTime;
    updatedAt: utcDateTime;
    accounts?: Accounts;
  }

  const PromoteUserRequestExample = #{ password: "password" };
  model PromoteUserRequest {
    password: string;
  }

  const AccountsExample = #{
    github_user_id: "yashikota",
    discord_channel_id: "9419543265265028",
    slack_channel_id: "HG824ORHQ250",
    qiita_user_id: "yashikota",
    zenn_username: "yashikota",
    wakatime_token: "yashikota",
  };
  model Accounts {
    github_user_id?: string;
    discord_channel_id?: string;
    slack_channel_id?: string;
    qiita_user_id?: string;
    zenn_username?: string;
    wakatime_token?: string;
  }

  model UpdateUserRequest {
    user_name?: UserID;
    email?: MailAddress;
    password?: Password;
    accounts?: Accounts;
  }

  @route("/api/v1/users/me")
  @tag("Users")
  @useAuth(BearerAuth)
  interface Users {
    @get
    @summary("Get user information")
    get(): User | ErrorResponse;

    @put
    @summary("Update user information")
    update(@body requestBody: UpdateUserRequest): User | ErrorResponse;

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
      @body requestBody: PromoteUserRequest,
    ): SuccessResponse | ErrorResponse;
  }
}
