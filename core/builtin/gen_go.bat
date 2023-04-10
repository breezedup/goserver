@echo off
set work_path=%cd%
set proto_path=%cd%\protocol
set protoc3=%cd%\..\..\bin\protoc-3.5.1-win32\bin\protoc.exe
set protoc-gen-go-plugin-path="%cd%\..\..\bin\protoc-gen-go.exe"

cd %proto_path%
 	for %%b in (,*.proto) do (
 	    echo %%b
	    %protoc3% --plugin=protoc-gen-go=%protoc-gen-go-plugin-path% --go_out=. %%b
	)
	cd ..
pause