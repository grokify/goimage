# PDFium

## Installation

The following is from: [https://jpadilla.com/2022/09/02/installing-pdfium-in-mac-os/](https://jpadilla.com/2022/09/02/installing-pdfium-in-mac-os/).

TIL: How to install pdfium in Mac OS using the pre-compiled binaries provided by https://github.com/bblanchon/pdfium-binaries.

```
$ mkdir ./pdfium
$ curl -sLo pdfium.tgz https://github.com/bblanchon/pdfium-binaries/releases/latest/download/pdfium-mac-x64.tgz
$ tar -xzvf pdfium.tgz -C ./pdfium
$ mv ./pdfium /usr/local/opt/
$ VERSION="$(cat /usr/local/opt/pdfium/VERSION | grep BUILD | sed -e 's/BUILD=//')" \
cat > /usr/local/lib/pkgconfig/pdfium.pc<< EOF
prefix=/usr/local/opt/pdfium
libdir=/usr/local/opt/pdfium/lib
includedir=/usr/local/opt/pdfium/include

Name: PDFium
Description: PDFium
Version: $VERSION
Requires:

Libs: -L\${libdir} -lpdfium
Cflags: -I\${includedir}
EOF
$ ln -s /usr/local/opt/pdfium/lib/libpdfium.dylib /usr/local/lib/libpdfium.dylib
```

## Troubleshooting

### Check to see if PDFium is findable by `pkg-config`

```
% pkg-config --cflags --libs pdfium
Package pdfium was not found in the pkg-config search path.
Perhaps you should add the directory containing `pdfium.pc'
to the PKG_CONFIG_PATH environment variable
No package 'pdfium' found
```

### Check `pkg-config` filepaths:

```
% pkg-config --variable pc_path pkg-config
```