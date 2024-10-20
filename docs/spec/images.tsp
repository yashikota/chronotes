using TypeSpec.Http;

namespace Chronotes {
  model UploadImageResponse {
    @doc("画像パス")
    @example("test123/01JAKVW7B6SRJZTV4AXV3E1FK1_my_dango.png")
    object_name: string;
  }

  @route("/api/v1/images")
  @tag("Images")
  @useAuth(BearerAuth)
  interface Images {
    @doc("対応画像形式はpngとjpegとwebpのみ。ファイルサイズは10MBまで")
    @post
    @summary("Upload an image")
    uploadImage(
      @header contentType: "multipart/form-data",
      image: bytes,
    ): UploadImageResponse | ErrorResponse;

    @doc("`object_name` を指定して画像を取得。URLは60秒間のみ有効")
    @get
    @summary("Get an image")
    getImage(@query object_name: string): bytes | ErrorResponse;

    @delete
    @summary("Delete an image")
    @doc("`object_name` を指定して削除")
    deleteImage(@query object_name: string): SuccessResponse | ErrorResponse;
  }
}
