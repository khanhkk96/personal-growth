cd ./di
del app_di.go
echo Deleted old generated file.
cd ./core
wire
move wire_gen.go ../app_di.go
echo Wire generation completed successfully.