# Sign Apps
Agar file executable (exe) yang dihasilkan dari kode Golang dikenali oleh Windows, pastikan Anda mengikuti langkah-langkah berikut:

1. **Compile Program untuk Windows**: Pastikan Anda mengompilasi program Golang Anda dengan target sistem operasi Windows. Jika Anda menggunakan sistem operasi selain Windows, Anda bisa menggunakan fitur cross-compilation di Golang. Berikut adalah cara untuk melakukannya:

    ```sh
    GOOS=windows GOARCH=amd64 go build -o program.exe main.go
    ```

    Ini akan menghasilkan file `program.exe` yang dapat dijalankan di Windows.

2. **Menggunakan Ikon Khusus (Optional)**: Untuk memberikan tampilan yang lebih profesional, Anda mungkin ingin menambahkan ikon khusus ke file executable Anda. Ini bisa dilakukan menggunakan resource file dan alat seperti `rsrc`.

    - Instal `rsrc`:
      ```sh
      go install github.com/akavel/rsrc@latest
      ```

    - Buat file resource dengan ikon:
      ```sh
      rsrc -manifest ransomdemo.manifest -ico favicon.ico -o rsrc.syso
      ```

    - Tambahkan file resource ke proyek Anda dan compile:
      ```sh
      GOOS=windows GOARCH=amd64 go build -o ransomdemo.exe main.go
      ```

3. **Menandatangani File Exe (Optional)**: Untuk menghindari peringatan dari Windows tentang keamanan, Anda bisa menandatangani file executable Anda menggunakan sertifikat digital. Ini melibatkan penggunaan alat seperti `signtool.exe` yang disediakan oleh [Microsoft](https://developer.microsoft.com/en-us/windows/downloads/windows-sdk/):

    ```sh
    signtool sign /a /t http://timestamp.digicert.com /fd SHA256 /tr http://timestamp.digicert.com ransomdemo.exe
    ```

4. **Mengatur File Metadata (Optional)**: Anda bisa menambahkan metadata seperti versi, nama produk, dan informasi perusahaan ke executable Anda menggunakan `goversioninfo`.

    - Instal `goversioninfo`:
      ```sh
      go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo@latest
      ```

    - Buat file `versioninfo.json` dengan konten seperti ini:
      ```json
      {
          "CompanyName": "Your Company",
          "FileDescription": "Your File Description",
          "FileVersion": "1.0.0.0",
          "ProductVersion": "1.0.0.0",
          "ProductName": "Your Product Name",
          "IconFile": "youricon.ico"
      }
      ```

    - Generate resource file:
      ```sh
      goversioninfo -manifest ransomdemo.manifest -o rsrc.syso
      ```

    - Compile ulang program dengan resource file yang baru:
      ```sh
      GOOS=windows GOARCH=amd64 go build -o ransomdemo.exe main.go
      ```

Dengan mengikuti langkah-langkah di atas, executable Golang Anda akan lebih mudah dikenali dan dijalankan oleh sistem operasi Windows tanpa peringatan keamanan yang mengganggu.