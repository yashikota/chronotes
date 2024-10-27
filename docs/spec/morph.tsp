using TypeSpec.Http;

namespace Chronotes {
  @route("/api/v1/morph")
  @tag("Morph")
  @useAuth(BearerAuth)
  interface Morhpological {
    @get
    @doc("文字列を送ると形態素解析した結果を返す")
    @summary("Morhpological analysis")
    getMorph(@query sentence: string): Morph | ErrorResponse;
  }
}
