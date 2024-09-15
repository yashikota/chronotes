import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";

using TypeSpec.Http;
@service({
  title: "Chronotes API",
})
@server("https://chronotes.yashikota.com", "Single server endpoint")

namespace chronotes;
