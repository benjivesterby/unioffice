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

package cid ;import (_e "encoding/xml";_b "fmt";_dg "github.com/unidoc/unioffice";_bb "github.com/unidoc/unioffice/common/logger";_d "strconv";);type CT_CommentId struct{ParaIdAttr string ;DurableIdAttr string ;};

// Validate validates the CommentsIds and its children
func (_ga *CommentsIds )Validate ()error {return _ga .ValidateWithPath ("C\u006f\u006d\u006d\u0065\u006e\u0074\u0073\u0049\u0064\u0073");};

// ValidateWithPath validates the CT_CommentsIds and its children, prefixing error messages with path
func (_bff *CT_CommentsIds )ValidateWithPath (path string )error {for _gc ,_ef :=range _bff .CommentId {if _eg :=_ef .ValidateWithPath (_b .Sprintf ("\u0025\u0073/\u0043\u006f\u006dm\u0065\u006e\u0074\u0049\u0064\u005b\u0025\u0064\u005d",path ,_gc ));_eg !=nil {return _eg ;
};};return nil ;};func (_be *CT_CommentId )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u0063i\u003a\u0070\u0061\u0072\u0061\u0049d"},Value :_b .Sprintf ("\u0025\u0076",_be .ParaIdAttr )});
start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u0063\u0069\u003ad\u0075\u0072\u0061\u0062\u006c\u0065\u0049\u0064"},Value :_b .Sprintf ("\u0025\u0076",_be .DurableIdAttr )});e .EncodeToken (start );e .EncodeToken (_e .EndElement {Name :start .Name });
return nil ;};func NewDecimaldurableId ()*DecimaldurableId {_ceb :=&DecimaldurableId {};return _ceb };func (_gb *CT_CommentsIds )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {e .EncodeToken (start );if _gb .CommentId !=nil {_ea :=_e .StartElement {Name :_e .Name {Local :"\u0063\u0069\u003ac\u006f\u006d\u006d\u0065\u006e\u0074\u0049\u0064"}};
for _ ,_ac :=range _gb .CommentId {e .EncodeElement (_ac ,_ea );};};e .EncodeToken (_e .EndElement {Name :start .Name });return nil ;};func (_ed *CommentsIds )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_ed .CT_CommentsIds =*NewCT_CommentsIds ();
_acf :for {_ec ,_fd :=d .Token ();if _fd !=nil {return _fd ;};switch _dcd :=_ec .(type ){case _e .StartElement :switch _dcd .Name {case _e .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002ec\u006f\u006d\u002f\u006f\u0066fi\u0063\u0065\u002f\u0077\u006f\u0072\u0064\u002f\u0032\u0030\u0031\u0036\u002f\u0077\u006f\u0072\u0064\u006d\u006c\u002f\u0063\u0069\u0064",Local :"\u0063o\u006d\u006d\u0065\u006e\u0074\u0049d"}:_de :=NewCT_CommentId ();
if _gbg :=d .DecodeElement (_de ,&_dcd );_gbg !=nil {return _gbg ;};_ed .CommentId =append (_ed .CommentId ,_de );default:_bb .Log .Debug ("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u006f\u006d\u006d\u0065\u006e\u0074\u0073\u0049\u0064\u0073\u0020\u0025\u0076",_dcd .Name );
if _gdd :=d .Skip ();_gdd !=nil {return _gdd ;};};case _e .EndElement :break _acf ;case _e .CharData :};};return nil ;};type CT_CommentsIds struct{CommentId []*CT_CommentId ;};

// ValidateWithPath validates the DecimaldurableId and its children, prefixing error messages with path
func (_bge *DecimaldurableId )ValidateWithPath (path string )error {return nil };func (_cae *DecimaldurableId )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {for _ ,_fe :=range start .Attr {if _fe .Name .Local =="\u0064u\u0072\u0061\u0062\u006c\u0065\u0049d"{_abf ,_fa :=_d .ParseInt (_fe .Value ,10,64);
if _fa !=nil {return _fa ;};_cae .DurableIdAttr =&_abf ;continue ;};};for {_fc ,_fea :=d .Token ();if _fea !=nil {return _b .Errorf ("\u0070\u0061\u0072\u0073i\u006e\u0067\u0020\u0044\u0065\u0063\u0069\u006d\u0061\u006cd\u0075r\u0061\u0062\u006c\u0065\u0049\u0064\u003a \u0025\u0073",_fea );
};if _aab ,_cab :=_fc .(_e .EndElement );_cab &&_aab .Name ==start .Name {break ;};};return nil ;};type CommentsIds struct{CT_CommentsIds };func NewCommentsIds ()*CommentsIds {_ebd :=&CommentsIds {};_ebd .CT_CommentsIds =*NewCT_CommentsIds ();return _ebd ;
};func (_ggc *DecimaldurableId )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {if _ggc .DurableIdAttr !=nil {start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u0063\u0069\u003ad\u0075\u0072\u0061\u0062\u006c\u0065\u0049\u0064"},Value :_b .Sprintf ("\u0025\u0076",*_ggc .DurableIdAttr )});
};start .Name .Local ="\u0063\u0069\u003a\u0064ec\u0069\u006d\u0061\u006c\u0064\u0075\u0072\u0061\u0062\u006c\u0065\u0049\u0064";return nil ;};

// ValidateWithPath validates the CommentsIds and its children, prefixing error messages with path
func (_gge *CommentsIds )ValidateWithPath (path string )error {if _ag :=_gge .CT_CommentsIds .ValidateWithPath (path );_ag !=nil {return _ag ;};return nil ;};func (_cf *CT_CommentsIds )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_dc :for {_cfa ,_beg :=d .Token ();
if _beg !=nil {return _beg ;};switch _ee :=_cfa .(type ){case _e .StartElement :switch _ee .Name {case _e .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002ec\u006f\u006d\u002f\u006f\u0066fi\u0063\u0065\u002f\u0077\u006f\u0072\u0064\u002f\u0032\u0030\u0031\u0036\u002f\u0077\u006f\u0072\u0064\u006d\u006c\u002f\u0063\u0069\u0064",Local :"\u0063o\u006d\u006d\u0065\u006e\u0074\u0049d"}:_dgd :=NewCT_CommentId ();
if _gg :=d .DecodeElement (_dgd ,&_ee );_gg !=nil {return _gg ;};_cf .CommentId =append (_cf .CommentId ,_dgd );default:_bb .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0043\u006f\u006d\u006d\u0065\u006e\u0074s\u0049d\u0073\u0020\u0025\u0076",_ee .Name );
if _ad :=d .Skip ();_ad !=nil {return _ad ;};};case _e .EndElement :break _dc ;case _e .CharData :};};return nil ;};type DecimaldurableId struct{DurableIdAttr *int64 ;};func (_adg *CommentsIds )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002ec\u006f\u006d\u002f\u006f\u0066fi\u0063\u0065\u002f\u0077\u006f\u0072\u0064\u002f\u0032\u0030\u0031\u0036\u002f\u0077\u006f\u0072\u0064\u006d\u006c\u002f\u0063\u0069\u0064"});
start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u0078\u006d\u006c\u006e\u0073\u003a\u0063\u0069"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002ec\u006f\u006d\u002f\u006f\u0066fi\u0063\u0065\u002f\u0077\u006f\u0072\u0064\u002f\u0032\u0030\u0031\u0036\u002f\u0077\u006f\u0072\u0064\u006d\u006c\u002f\u0063\u0069\u0064"});
start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0077"},Value :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063\u0065s\u0073i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u00306\u002fm\u0061\u0069n"});
start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});
start .Name .Local ="\u0063\u0069\u003a\u0063\u006f\u006d\u006d\u0065\u006et\u0073\u0049\u0064\u0073";return _adg .CT_CommentsIds .MarshalXML (e ,start );};func NewCT_CommentId ()*CT_CommentId {_g :=&CT_CommentId {};return _g };func (_a *CT_CommentId )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {for _ ,_c :=range start .Attr {if _c .Name .Local =="\u0070\u0061\u0072\u0061\u0049\u0064"{_bbb ,_cg :=_c .Value ,error (nil );
if _cg !=nil {return _cg ;};_a .ParaIdAttr =_bbb ;continue ;};if _c .Name .Local =="\u0064u\u0072\u0061\u0062\u006c\u0065\u0049d"{_bf ,_ce :=_c .Value ,error (nil );if _ce !=nil {return _ce ;};_a .DurableIdAttr =_bf ;continue ;};};for {_gd ,_bd :=d .Token ();
if _bd !=nil {return _b .Errorf ("\u0070a\u0072\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005f\u0043\u006fm\u006d\u0065\u006e\u0074\u0049\u0064\u003a\u0020\u0025\u0073",_bd );};if _ca ,_eb :=_gd .(_e .EndElement );_eb &&_ca .Name ==start .Name {break ;
};};return nil ;};

// Validate validates the CT_CommentsIds and its children
func (_ab *CT_CommentsIds )Validate ()error {return _ab .ValidateWithPath ("\u0043\u0054\u005f\u0043\u006f\u006d\u006d\u0065\u006et\u0073\u0049\u0064\u0073");};

// Validate validates the DecimaldurableId and its children
func (_fec *DecimaldurableId )Validate ()error {return _fec .ValidateWithPath ("\u0044\u0065c\u0069\u006d\u0061l\u0064\u0075\u0072\u0061\u0062\u006c\u0065\u0049\u0064");};

// Validate validates the CT_CommentId and its children
func (_cga *CT_CommentId )Validate ()error {return _cga .ValidateWithPath ("\u0043\u0054\u005fC\u006f\u006d\u006d\u0065\u006e\u0074\u0049\u0064");};func NewCT_CommentsIds ()*CT_CommentsIds {_bg :=&CT_CommentsIds {};return _bg };

// ValidateWithPath validates the CT_CommentId and its children, prefixing error messages with path
func (_aa *CT_CommentId )ValidateWithPath (path string )error {return nil };func init (){_dg .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002ec\u006f\u006d\u002f\u006f\u0066fi\u0063\u0065\u002f\u0077\u006f\u0072\u0064\u002f\u0032\u0030\u0031\u0036\u002f\u0077\u006f\u0072\u0064\u006d\u006c\u002f\u0063\u0069\u0064","\u0043\u0054\u005f\u0043\u006f\u006d\u006d\u0065\u006et\u0073\u0049\u0064\u0073",NewCT_CommentsIds );
_dg .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002ec\u006f\u006d\u002f\u006f\u0066fi\u0063\u0065\u002f\u0077\u006f\u0072\u0064\u002f\u0032\u0030\u0031\u0036\u002f\u0077\u006f\u0072\u0064\u006d\u006c\u002f\u0063\u0069\u0064","\u0043\u0054\u005fC\u006f\u006d\u006d\u0065\u006e\u0074\u0049\u0064",NewCT_CommentId );
_dg .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002ec\u006f\u006d\u002f\u006f\u0066fi\u0063\u0065\u002f\u0077\u006f\u0072\u0064\u002f\u0032\u0030\u0031\u0036\u002f\u0077\u006f\u0072\u0064\u006d\u006c\u002f\u0063\u0069\u0064","c\u006f\u006d\u006d\u0065\u006e\u0074\u0073\u0049\u0064\u0073",NewCommentsIds );
_dg .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002ec\u006f\u006d\u002f\u006f\u0066fi\u0063\u0065\u002f\u0077\u006f\u0072\u0064\u002f\u0032\u0030\u0031\u0036\u002f\u0077\u006f\u0072\u0064\u006d\u006c\u002f\u0063\u0069\u0064","\u0064\u0065c\u0069\u006d\u0061l\u0064\u0075\u0072\u0061\u0062\u006c\u0065\u0049\u0064",NewDecimaldurableId );
};