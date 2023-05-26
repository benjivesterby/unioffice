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

package algo ;import _f "strconv";func RepeatString (s string ,cnt int )string {if cnt <=0{return "";};_ffg :=make ([]byte ,len (s )*cnt );_gbf :=[]byte (s );for _c :=0;_c < cnt ;_c ++{copy (_ffg [_c :],_gbf );};return string (_ffg );};func _g (_ff byte )bool {return _ff >='0'&&_ff <='9'};

// NaturalLess compares two strings in a human manner so rId2 sorts less than rId10
func NaturalLess (lhs ,rhs string )bool {_d ,_dc :=0,0;for _d < len (lhs )&&_dc < len (rhs ){_dg :=lhs [_d ];_db :=rhs [_dc ];_fg :=_g (_dg );_gb :=_g (_db );switch {case _fg &&!_gb :return true ;case !_fg &&_gb :return false ;case !_fg &&!_gb :if _dg !=_db {return _dg < _db ;};_d ++;_dc ++;default:_e :=_d +1;_b :=_dc +1;for _e < len (lhs )&&_g (lhs [_e ]){_e ++;};for _b < len (rhs )&&_g (rhs [_b ]){_b ++;};_bb ,_ :=_f .ParseUint (lhs [_d :_e ],10,64);_bg ,_ :=_f .ParseUint (rhs [_d :_b ],10,64);if _bb !=_bg {return _bb < _bg ;};_d =_e ;_dc =_b ;};};return len (lhs )< len (rhs );};