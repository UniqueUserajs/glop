g++ -o glop.o -m32 -c -Iinclude glop.cpp
g++ -m32 -shared -dynamiclib -Wp,1 -o libglop.so glop.o
rm -f glop.o
mkdir -p lib
mv libglop.so lib
