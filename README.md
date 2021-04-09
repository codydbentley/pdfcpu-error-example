### pdfcpu error example

For use in helping debug [https://github.com/pdfcpu/pdfcpu/issues/320](https://github.com/pdfcpu/pdfcpu/issues/320)

To Run:
```
go run main.go
```

System information:
```
OS: Windows 10 64-bit
Golang: go1.15.2 windows/amd64
pdfcpu: v0.3.11
```

Task details:
- merge `page1.pdf` and `page2.pdf` using `api.Merge()`
- apply an image watermark to `input.pdf` using `api.AddWatermarks()`

#### Desired Output:
2-page PDF called `output.pdf` with the text hidden
behind a white image watermark on page 1

#### Success Case:
With Windows Date & Time Settings set to Time Zone 
`(UTC -07:00) Mountain Time (US & Canada)`, the program functions as expected.

#### Failure Case:
With Windows Date & Time Setting set to Time Zone
`(UTC -10:00) Hawaii`, the program fails with the following error:
```
add watermark error: pdfcpu: validateDateObject: <D:20210409072606+-10'00'> invalid date
```

#### Stack Trace:
```
Variables:
o = {github.com/pdfcpu/pdfcpu/pkg/pdfcpu.StringLiteral}(D:20210409073107+-10'00')

github.com/pdfcpu/pdfcpu/pkg/pdfcpu/validate.validateDateObject at objects.go:183
github.com/pdfcpu/pdfcpu/pkg/pdfcpu/validate.validateInfoDictDate at info.go:48
github.com/pdfcpu/pdfcpu/pkg/pdfcpu/validate.validateDocInfoDictEntry at info.go:123
github.com/pdfcpu/pdfcpu/pkg/pdfcpu/validate.validateDocumentInfoDict at info.go:153
github.com/pdfcpu/pdfcpu/pkg/pdfcpu/validate.validateDocumentInfoObject at info.go:182
github.com/pdfcpu/pdfcpu/pkg/pdfcpu/validate.XRefTable at xReftable.go:46
github.com/pdfcpu/pdfcpu/pkg/api.readAndValidate at api.go:111
github.com/pdfcpu/pdfcpu/pkg/api.readValidateAndOptimize at api.go:121
github.com/pdfcpu/pdfcpu/pkg/api.AddWatermarks at stamp.go:220
main.main at main.go:66
runtime.main at proc.go:204
runtime.goexit at asm_amd64.s:1374
 - Async Stack Trace
runtime.rt0_go at asm_amd64.s:220
```