//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package tempstorage ;import _a "io";

// RemoveAll removes all files according to the dir argument prefix.
func RemoveAll (dir string )error {return _be .RemoveAll (dir )};

// File is a representation of a storage file
// with Read, Write, Close and Name methods identical to os.File.
type File interface{_a .Reader ;_a .ReaderAt ;_a .Writer ;_a .Closer ;Name ()string ;};

// Add reads a file from a disk and adds it to the storage.
func Add (path string )error {return _be .Add (path )};

// TempDir creates a name for a new temp directory using a pattern argument.
func TempDir (pattern string )(string ,error ){return _be .TempDir (pattern )};

// SetAsStorage changes temporary storage to newStorage.
func SetAsStorage (newStorage storage ){_be =newStorage };type storage interface{Open (_d string )(File ,error );TempFile (_f ,_aa string )(File ,error );TempDir (_c string )(string ,error );RemoveAll (_e string )error ;Add (_g string )error ;};

// Open returns tempstorage File object by name.
func Open (path string )(File ,error ){return _be .Open (path )};var _be storage ;

// TempFile creates new empty file in the storage and returns it.
func TempFile (dir ,pattern string )(File ,error ){return _be .TempFile (dir ,pattern )};