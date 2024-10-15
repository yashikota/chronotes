import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";
import "@typespec/versioning";

import "./auth.tsp";
import "./users.tsp";
import "./notes.tsp";
import "./images.tsp";
import "./admin.tsp";

using TypeSpec.Http;
using TypeSpec.Rest;

@service({
  title: "Chronotes API",
  description: "Chronotes API specification",
})
@Versioning.versioned(ApiVersion)
@server("http://localhost:5678", "dev")
@server("https://chronotes.yashikota.com", "prd")
namespace Chronotes {
  enum ApiVersion {
    v1,
  }

  // misc
  const SuccessResponseExample = #{ message: "Success" };
  @example(SuccessResponseExample)
  model SuccessResponse {
    message: string;
  }

  const ErrorResponseExample = #{ message: "Error" };
  @example(ErrorResponseExample)
  @error
  model ErrorResponse {
    message: string;
  }

  @route("/api/v1/health")
  @get
  @summary("Check API health")
  @tag("Health")
  @useAuth(NoAuth)
  op healthCheck(): SuccessResponse | ErrorResponse;
}
