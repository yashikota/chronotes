import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";
import "@typespec/versioning";

import "./model.tsp";
import "./auth.tsp";
import "./users.tsp";
import "./notes.tsp";
import "./images.tsp";
import "./admin.tsp";
import "./morph.tsp";

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
  model SuccessResponse {
    @doc("Success message")
    @example("Success")
    message: string;
  }

  @error
  model ErrorResponse {
    @doc("Error message")
    @example("Error")
    message: string;
  }

  @route("/api/v1/health")
  @get
  @summary("Check API health")
  @tag("Health")
  @useAuth(NoAuth)
  op healthCheck(): SuccessResponse | ErrorResponse;
}
