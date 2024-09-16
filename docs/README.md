# docs

## TypeSpec

APIのスキーマ定義に `TypeSpec(tsp)` を用いている。  

1. tspのinstall

    ```sh
    npm install -g @typespec/compiler
    ```

2. ディレクトリの移動

    ```sh
    cd docs/spec
    ```

3. tspの依存関係のinstall

    ```sh
    tsp install
    ```

4. tspのcompile

    ```sh
    tsp compile .
    ```

なお、4の手順のみ `task tsp` で行える。  
ただし、事前に3までの手順を行う必要がある。  

## API

OpenAPIのコーディング規約は以下の記事を参考にしている。  

<https://future-architect.github.io/coding-standards/documents/forOpenAPISpecification/OpenAPI_Specification_3.0.3.html>
