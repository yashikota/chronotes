using TypeSpec.Http;

namespace Chronotes {
  model UploadImageResponse {
    @doc("変換後の画像パス")
    @example("img/01JA0QTJB2MHTBNJV3SJ9KYN9J.jpg")
    converted: string;

    @doc("変換前の画像パス")
    @example("img/01JA0QTJB2MHTBNJV3SJ9KYN9J_original.jpg")
    original: string;
  }

  @route("/api/v1/images")
  @tag("Images")
  @useAuth(BearerAuth)
  interface Images {
    @post
    @summary("Upload an image")
    @doc("対応画像形式はpngとjpegとwebpのみ。ファイルサイズは10MBまで。")
    uploadImage(
      @header contentType: "multipart/form-data",
      image: bytes,
    ): UploadImageResponse | ErrorResponse;

    @delete
    @summary("Delete an image")
    @doc("`image_path`には画像のIDを指定して削除")
    deleteImage(@query image_id: string): SuccessResponse | ErrorResponse;
  }
}
