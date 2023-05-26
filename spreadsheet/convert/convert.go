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

package convert ;import (_c "github.com/unidoc/unioffice/common/logger";_ef "github.com/unidoc/unioffice/common/tempstorage";_ba "github.com/unidoc/unioffice/internal/convertutils";_e "github.com/unidoc/unioffice/measurement";_gd "github.com/unidoc/unioffice/schema/soo/dml";_a "github.com/unidoc/unioffice/schema/soo/dml/chart";_dd "github.com/unidoc/unioffice/schema/soo/ofc/sharedTypes";_bc "github.com/unidoc/unioffice/schema/soo/sml";_cd "github.com/unidoc/unioffice/spreadsheet";_ag "github.com/unidoc/unioffice/spreadsheet/reference";_df "github.com/unidoc/unipdf/v3/creator";_bb "github.com/unidoc/unipdf/v3/model";_g "image";_d "strconv";);type colWidthRange struct{_cdba int ;_ebdee int ;_feeg float64 ;_cgaf *style ;};type rowInfo struct{_ecde float64 ;_bfdc bool ;_ccdb float64 ;_adaaf *style ;_afgd []*cell ;_fec float64 ;};func (_cbb *convertContext )makePages (){for _ ,_fe :=range _cbb ._gffe {for _ ,_fae :=range _cbb ._cegg {_fe ._fce =append (_fe ._fce ,&page {_dbaa :[]*pageRow {},_debe :_fe ,_fbde :_fae });};};};func (_abg *convertContext )getStyleFromRPrElt (_cfbf *_bc .CT_RPrElt )*style {if _cfbf ==nil {return nil ;};_fdb :=&style {};_fdb ._eade =&_cfbf .RFont .ValAttr ;if _afef :=_cfbf .B ;_afef !=nil {_cfc :=_afef .ValAttr ==nil ||*_afef .ValAttr ;_fdb ._cedd =&_cfc ;};if _ggbg :=_cfbf .I ;_ggbg !=nil {_edbb :=_ggbg .ValAttr ==nil ||*_ggbg .ValAttr ;_fdb ._ecdf =&_edbb ;};if _cbadg :=_cfbf .U ;_cbadg !=nil {_gge :=_cbadg .ValAttr ==_bc .ST_UnderlineValuesSingle ||_cbadg .ValAttr ==_bc .ST_UnderlineValuesUnset ;_fdb ._ggbf =&_gge ;};if _eege :=_cfbf .VertAlign ;_eege !=nil {_dec :=_eege .ValAttr ==_dd .ST_VerticalAlignRunSuperscript ;_fdb ._cfff =&_dec ;_gaec :=_eege .ValAttr ==_dd .ST_VerticalAlignRunSubscript ;_fdb ._faa =&_gaec ;};if _faac :=_cfbf .Sz ;_faac !=nil {_ace :=_faac .ValAttr /12*_ba .DefaultFontSize ;_fdb ._gaedc =&_ace ;};if _deca :=_cfbf .Color ;_deca !=nil {_fdb ._adfb =_abg .getColorStringFromSmlColor (_deca );};return _fdb ;};type cell struct{_gab _bc .ST_CellType ;_bbeb int ;_accd float64 ;_ffce []*line ;_bacg float64 ;_cbad float64 ;_daea float64 ;_gdag float64 ;_agf float64 ;_acbe *_df .TextStyle ;_abe *border ;_fged *border ;_dgg *border ;_bfcd *border ;_bagc bool ;_dbe bool ;};const _ce =0.25;const _fb =2;type page struct{_dbaa []*pageRow ;_edf bool ;_abbg []*_df .Image ;_debe *pagespan ;_fbde *rowspan ;};func (_bfd *convertContext )getImage (_cef _g .Image ,_bfa ,_ceg ,_cff ,_ccb ,_efa ,_cfbb float64 ,_fffg _ba .ImgPart )*_df .Image {_ccb +=_bfd ._acdc ;_cff +=_bfd ._efae ;_bbeg ,_babc :=_ba .GetImage (_bfd ._cbaf ,_cef ,_bfa ,_ceg ,_cff ,_ccb ,_efa ,_cfbb ,_fffg );if _babc !=nil {_c .Log .Debug ("\u0043\u0061\u006eno\u0074\u0020\u0067\u0065\u0074\u0020\u0061\u006e\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073",_babc );return nil ;};return _bbeg ;};

// FontStyle represents a kind of font styling. It can be FontStyle_Regular, FontStyle_Bold, FontStyle_Italic and FontStyle_BoldItalic.
type FontStyle =_ba .FontStyle ;var _ebf =3.025/_cac (1);func (_abb *convertContext )makeCells (){_dde :=_abb ._ccbd ;_dad :=_dde .Rows ();_be :=0;for _ ,_bff :=range _abb ._gag {_bff ._afgd =[]*cell {};_caf :=0.0;_dbc :=_bff ._adaaf ;if _bff ._bfdc {_fbb :=_dad [_be ];_be ++;_bdf :=_bff ._ccdb ;for _ ,_cg :=range _fbb .Cells (){_bce ,_dgc :=_ag .ParseCellReference (_cg .Reference ());if _dgc !=nil {_c .Log .Debug ("\u0043\u0061\u006e\u006eo\u0074\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0061\u0020r\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u003a \u0025\u0073",_dgc );continue ;};_gaf :=_abb ._dgbc [_bce .ColumnIdx ];_cca :=_gaf ._bec ;_cge :=_cca ;_dgf :=_gaf ._bdaa ;var _fab ,_dbf ,_adac ,_aec bool ;for _ ,_bbde :=range _abb ._eacd {if _bce .RowIdx >=_bbde ._aecb &&_bce .RowIdx <=_bbde ._aee &&_bce .ColumnIdx >=_bbde ._becb &&_bce .ColumnIdx <=_bbde ._bfcg {if _bce .ColumnIdx ==_bbde ._becb &&_bce .RowIdx ==_bbde ._aecb {_cca =_bbde ._gfeb ;_bdf =_bbde ._ebdec ;};_fab =_bce .RowIdx !=_bbde ._aecb ;_dbf =_bce .RowIdx !=_bbde ._aee ;_adac =_bce .ColumnIdx !=_bbde ._becb ;_aec =_bce .ColumnIdx !=_bbde ._bfcg ;};};_gcc :=_abb .getStyleFromCell (_cg ,_dbc ,_dgf );var _gcb ,_ff ,_ebd bool ;var _eeb ,_ffe ,_gde ,_dgfd *border ;var _dcad _bc .ST_VerticalAlignment ;var _gacb _bc .ST_HorizontalAlignment ;if _gcc !=nil {if !_fab {_eeb =_gcc ._bfbb ;};if !_dbf {_ffe =_gcc ._gece ;};if !_adac {_gde =_gcc ._gdfe ;};if !_aec {_dgfd =_gcc ._bag ;};if _ffe !=nil &&_ffe ._eeg > _caf {_caf =_ffe ._eeg ;};_dcad =_gcc ._geg ;_gacb =_gcc ._bebc ;if _gcc ._cfff !=nil {_gcb =*_gcc ._cfff ;};if _gcc ._faa !=nil {_ff =*_gcc ._faa ;};_ebd =_gcc ._afeb ;};_ggg ,_baf :=_abb .getContentFromCell (_cg ,_gcc ,_cca ,_ebd );_abf :=&cell {_gab :_baf ,_bacg :_cca ,_cbad :_cge ,_daea :_bdf ,_ffce :_ggg ,_abe :_eeb ,_fged :_ffe ,_dgg :_gde ,_bfcd :_dgfd ,_bagc :_gcb ,_dbe :_ff };_abb .alignSymbolsHorizontally (_abf ,_gacb );_abb .alignSymbolsVertically (_abf ,_dcad );_bff ._afgd =append (_bff ._afgd ,_abf );};};_bff ._fec =_caf ;};};func (_dgb *convertContext )drawSheet (){for _dedaa ,_eab :=range _dgb ._gffe {_gdb :=len (_eab ._fce );if _dedaa ==len (_dgb ._gffe )-1{for _ffg :=len (_eab ._fce )-1;_ffg >=0;_ffg --{if !_eab ._fce [_ffg ]._edf {_gdb =_ffg ;};};};_gbf :=_eab ._fce [:_gdb ];for _ ,_dbfc :=range _gbf {_dgb ._cbaf .NewPage ();_dgb .drawPage (_dbfc );};};};func (_ffcd *convertContext )imageFromAnchor (_eec *anchor ,_ecgb ,_fcb float64 )_g .Image {if _eec ._gbgd !=nil {return _eec ._gbgd ;};if _eec ._dggb !=nil {_dff ,_dcda :=_ba .MakeImageFromChartSpace (_eec ._dggb ,_ecgb ,_fcb ,_ffcd ._ffb );if _dcda !=nil {_c .Log .Debug ("C\u0061\u006e\u006e\u006f\u0074\u0020\u006d\u0061\u006b\u0065\u0020\u0061\u006e\u0020\u0069\u006d\u0061\u0067e\u0020\u0066\u0072\u006f\u006d\u0020\u0063\u0068\u0061\u0072tS\u0070\u0061\u0063e\u003a \u0025\u0073",_dcda );return nil ;};return _dff ;};return nil ;};func (_gbd *convertContext )makeCols (){_gbdb :=_gbd ._ccbd ;_gf :=_gbdb .X ();_ebc :=[]*colInfo {};_ega :=[]colWidthRange {};if _bbd :=_gf .Cols ;len (_bbd )> 0{for _ ,_cbc :=range _bbd [0].Col {_dfaf :=65.0;if _ab :=_cbc .WidthAttr ;_ab !=nil {if *_ab > 0.83{*_ab -=0.83;};if *_ab <=1{_dfaf =*_ab *11;}else {_dfaf =5+*_ab *6;};};_adc :=int (_cbc .MinAttr -1);_dca :=int (_cbc .MaxAttr -1);_ega =append (_ega ,colWidthRange {_cdba :_adc ,_ebdee :_dca ,_feeg :_dfaf ,_cgaf :_gbd .getStyle (_cbc .StyleAttr )});};};_gfa :=0;for _cdf :=0;_cdf <=_gbd ._fgdb ;_cdf ++{var _gac float64 ;var _ecd *style ;if _gfa >=len (_ega ){_gac =65;}else {_cf :=_ega [_gfa ];if _cdf >=_cf ._cdba &&_cdf <=_cf ._ebdee {_gac =_cf ._feeg ;_ecd =_cf ._cgaf ;if _cdf ==_cf ._ebdee {_gfa ++;};}else {_gac =65;};};_ebc =append (_ebc ,&colInfo {_bec :_gac ,_bdaa :_ecd });};_gbd ._dgbc =_ebc ;};func (_gfbg *convertContext )getColorStringFromSmlColor (_acgf *_bc .CT_Color )*string {var _aag string ;if _acgf .RgbAttr !=nil {_aag =*_acgf .RgbAttr ;}else if _acgf .IndexedAttr !=nil &&*_acgf .IndexedAttr < 64{_aag =_acfe [*_acgf .IndexedAttr ];}else if _acgf .ThemeAttr !=nil {_cda :=*_acgf .ThemeAttr ;_aag =_gfbg .getColorFromTheme (_cda );};if _aag ==""{return nil ;};if len (_aag )> 6{_aag =_aag [(len (_aag )-6):];};if _acgf .TintAttr !=nil {_dddd :=*_acgf .TintAttr ;_aag =_ba .AdjustColorByTint (_aag ,_dddd );};_aag ="\u0023"+_aag ;return &_aag ;};const _eb =1.5;func (_fcc *convertContext )distributeAnchors (){for _ ,_dae :=range _fcc ._bdb {_fed ,_bgbag :=_dae ._acga ,_dae ._agdg ;_gbdg ,_cdc :=_dae ._gdd ,_dae ._caba ;_gff ,_bbda :=_dae ._cgf ,_dae ._fd ;_gef ,_cec :=_dae ._dbgc ,_dae ._fbcb ;var _fgba ,_fbca ,_feg ,_dba *page ;for _ ,_adcg :=range _fcc ._gffe {for _ ,_ebde :=range _adcg ._fce {if _fed >=_ebde ._fbde ._ccagg &&_fed < _ebde ._fbde ._geb {if _gbdg >=_ebde ._debe ._gead &&_gbdg < _ebde ._debe ._bgbad {_ebde ._edf =true ;_fgba =_ebde ;};if _gef >=_ebde ._debe ._gead &&_gef < _ebde ._debe ._bgbad {_ebde ._edf =true ;_fbca =_ebde ;};};if _gff >=_ebde ._fbde ._ccagg &&_gff < _ebde ._fbde ._geb {if _gbdg >=_ebde ._debe ._gead &&_gbdg < _ebde ._debe ._bgbad {_ebde ._edf =true ;_dba =_ebde ;};if _gef >=_ebde ._debe ._gead &&_gef < _ebde ._debe ._bgbad {_ebde ._edf =true ;_feg =_ebde ;};};};};_bad :=_fgba !=_fbca ;_bef :=_fgba !=_dba ;if _bad &&_bef {_aaf :=_fcc ._dgbc [_gbdg ]._gccb +_e .FromEMU (_cdc );_fafg :=_fgba ._debe ._aggb ;_deb :=_fcc ._dgbc [_gef ]._gccb +_e .FromEMU (_cec );_cae :=_fcc ._gag [_fed ]._ecde +_e .FromEMU (_bgbag );_gaa :=float64 (_fgba ._fbde ._ebcd );_gdf :=_fcc ._gag [_gff ]._ecde +_e .FromEMU (_bbda );_ea :=_deb +_fafg -_aaf ;_gdfg :=_gdf +_gaa -_cae ;_cgg :=_fcc .imageFromAnchor (_dae ,_ea ,_gdfg );_fgba ._abbg =append (_fgba ._abbg ,_fcc .getImage (_cgg ,_gdfg ,_ea ,_aaf ,_cae ,_fafg -_aaf ,_gaa -_cae ,_ba .ImgPart_lt ));_fbca ._abbg =append (_fbca ._abbg ,_fcc .getImage (_cgg ,_gdfg ,_ea ,0,_cae ,_fafg -_aaf ,_gaa -_cae ,_ba .ImgPart_rt ));_dba ._abbg =append (_dba ._abbg ,_fcc .getImage (_cgg ,_gdfg ,_ea ,_aaf ,0,_fafg -_aaf ,_gaa -_cae ,_ba .ImgPart_lb ));_feg ._abbg =append (_feg ._abbg ,_fcc .getImage (_cgg ,_gdfg ,_ea ,0,0,_fafg -_aaf ,_gaa -_cae ,_ba .ImgPart_rb ));}else if _bad {_ccag :=_fcc ._gag [_fed ]._ecde +_e .FromEMU (_bgbag );_fac :=_fcc ._gag [_gff ]._ecde +_e .FromEMU (_bbda );_cab :=_fcc ._dgbc [_gbdg ]._gccb +_e .FromEMU (_cdc );_gea :=_fgba ._debe ._aggb ;_bbaf :=_fcc ._dgbc [_gef ]._gccb +_e .FromEMU (_cec );_cbe :=_bbaf +_gea -_cab ;_acgg :=_fac -_ccag ;_eefe :=_fcc .imageFromAnchor (_dae ,_cbe ,_acgg );_fgba ._abbg =append (_fgba ._abbg ,_fcc .getImage (_eefe ,_acgg ,_cbe ,_cab ,_ccag ,_gea -_cab ,0,_ba .ImgPart_l ));_fbca ._abbg =append (_fbca ._abbg ,_fcc .getImage (_eefe ,_acgg ,_cbe ,0,_ccag ,_gea -_cab ,0,_ba .ImgPart_r ));}else if _bef {_acb :=_fcc ._dgbc [_gbdg ]._gccb +_e .FromEMU (_cdc );_gafc :=_fcc ._dgbc [_gef ]._gccb +_e .FromEMU (_cec );_af :=_fcc ._gag [_fed ]._ecde +_e .FromEMU (_bgbag );_dfb :=float64 (_fgba ._fbde ._ebcd );_eefa :=_fcc ._gag [_gff ]._ecde +_e .FromEMU (_bbda );_bbb :=_gafc -_acb ;_dcd :=_eefa +_dfb -_af ;_bbe :=_fcc .imageFromAnchor (_dae ,_bbb ,_dcd );_fgba ._abbg =append (_fgba ._abbg ,_fcc .getImage (_bbe ,_dcd ,_bbb ,_acb ,_af ,0,_dfb -_af ,_ba .ImgPart_t ));_dba ._abbg =append (_dba ._abbg ,_fcc .getImage (_bbe ,_dcd ,_bbb ,_acb ,0,0,_dfb -_af ,_ba .ImgPart_b ));}else {_deda :=_fcc ._dgbc [_gbdg ]._gccb +_e .FromEMU (_cdc );_cga :=_fcc ._dgbc [_gef ]._gccb +_e .FromEMU (_cec );_dga :=_fcc ._gag [_fed ]._ecde +_e .FromEMU (_bgbag );_egad :=_fcc ._gag [_gff ]._ecde +_e .FromEMU (_bbda );_effb :=_cga -_deda ;_ddc :=_egad -_dga ;_effc :=_fcc .imageFromAnchor (_dae ,_effb ,_ddc );_fgba ._abbg =append (_fgba ._abbg ,_fcc .getImage (_effc ,_ddc ,_effb ,_deda ,_dga ,0,0,_ba .ImgPart_whole ));};};};func (_fga *convertContext )fillPages (){for _ege ,_gec :=range _fga ._cegg {_bfg :=_fga ._gag [_gec ._ccagg :_gec ._geb ];for _cfd ,_ddg :=range _bfg {_dac :=0;_agbb :=0.0;_dfd :=[]*cell {};if _ddg ._bfdc {for _ ,_faf :=range _ddg ._afgd {_gbc :=_fga ._gffe [_dac ];_fga ._dgbe =_gbc ._fce [_ege ];_fga ._dgbe ._edf =true ;_cgc :=_faf ._cbad ;if _agbb +_cgc > _gbc ._aggb {_fga .addRowToPage (_dfd ,_cfd );_dfd =[]*cell {_faf };_agbb =_cgc ;_dac ++;}else {_faf ._accd =_agbb ;_dfd =append (_dfd ,_faf );_agbb +=_cgc ;};};if len (_dfd )> 0{_aga :=_fga ._gffe [_dac ];_fga ._dgbe =_aga ._fce [_ege ];_fga ._dgbe ._edf =true ;_fga .addRowToPage (_dfd ,_cfd );};};};};};type pagespan struct{_aggb float64 ;_fce []*page ;_gead int ;_bgbad int ;};func _ebdf (_cdeb ,_ddbf *style ){if _ddbf ==nil {return ;};if _cdeb ==nil {_cdeb =_ddbf ;return ;};if _cdeb ._eade ==nil {_cdeb ._eade =_ddbf ._eade ;};if _cdeb ._adfb ==nil {_cdeb ._adfb =_ddbf ._adfb ;};if _cdeb ._gaedc ==nil {_cdeb ._gaedc =_ddbf ._gaedc ;};if _cdeb ._cedd ==nil {_cdeb ._cedd =_ddbf ._cedd ;};if _cdeb ._ecdf ==nil {_cdeb ._ecdf =_ddbf ._ecdf ;};if _cdeb ._ggbf ==nil {_cdeb ._ggbf =_ddbf ._ggbf ;};if _cdeb ._cfff ==nil {_cdeb ._cfff =_ddbf ._cfff ;};if _cdeb ._faa ==nil {_cdeb ._faa =_ddbf ._faa ;};if _cdeb ._bfbb ==nil {_cdeb ._bfbb =_ddbf ._bfbb ;};if _cdeb ._gece ==nil {_cdeb ._gece =_ddbf ._gece ;};if _cdeb ._gdfe ==nil {_cdeb ._gdfe =_ddbf ._gdfe ;};if _cdeb ._bag ==nil {_cdeb ._bag =_ddbf ._bag ;};if _cdeb ._geg ==_bc .ST_VerticalAlignmentUnset {_cdeb ._geg =_ddbf ._geg ;};if _cdeb ._bebc ==_bc .ST_HorizontalAlignmentUnset {_cdeb ._bebc =_ddbf ._bebc ;};};func (_ffc *convertContext )makeRowspans (){var _beb float64 ;_adaca :=0;for _gfb ,_fcg :=range _ffc ._gag {_bfb :=_fcg ._ccdb +_fcg ._fec ;if _beb +_bfb <=_ffc ._abaa {_fcg ._ecde =_beb ;_beb +=_bfb ;}else {_ffc ._cegg =append (_ffc ._cegg ,&rowspan {_ebcd :_beb ,_ccagg :_adaca ,_geb :_gfb });_adaca =_gfb ;_fcg ._ecde =0;_beb =_bfb ;};};_ffc ._cegg =append (_ffc ._cegg ,&rowspan {_ebcd :_beb ,_ccagg :_adaca ,_geb :len (_ffc ._gag )});};func _adb (_bgef *symbol ){_dgag :=_df .New ();_cee :=_dgag .NewStyledParagraph ();_cee .SetMargins (0,0,0,0);_egdg :=_cee .Append (_bgef ._eabb );if _bgef ._gee !=nil {_egdg .Style =*_bgef ._gee ;};_bgef ._agga =_cee .Height ();if _bgef ._aeab ==0{_bgef ._aeab =_cee .Width ();};};type mergedCell struct{_aecb uint32 ;_becb uint32 ;_aee uint32 ;_bfcg uint32 ;_gfeb float64 ;_ebdec float64 ;};var _acfe =[]string {"\u0030\u0030\u0030\u0030\u0030\u0030","\u0066\u0066\u0066\u0066\u0066\u0066","\u0066\u0066\u0030\u0030\u0030\u0030","\u0030\u0030\u0066\u0066\u0030\u0030","\u0030\u0030\u0030\u0030\u0066\u0066","\u0066\u0066\u0066\u0066\u0030\u0030","\u0066\u0066\u0030\u0030\u0066\u0066","\u0030\u0030\u0066\u0066\u0066\u0066","\u0030\u0030\u0030\u0030\u0030\u0030","\u0066\u0066\u0066\u0066\u0066\u0066","\u0066\u0066\u0030\u0030\u0030\u0030","\u0030\u0030\u0066\u0066\u0030\u0030","\u0030\u0030\u0030\u0030\u0066\u0066","\u0066\u0066\u0066\u0066\u0030\u0030","\u0066\u0066\u0030\u0030\u0066\u0066","\u0030\u0030\u0066\u0066\u0066\u0066","\u0038\u0030\u0030\u0030\u0030\u0030","\u0030\u0030\u0038\u0030\u0030\u0030","\u0030\u0030\u0030\u0030\u0038\u0030","\u0038\u0030\u0038\u0030\u0030\u0030","\u0038\u0030\u0030\u0030\u0038\u0030","\u0030\u0030\u0038\u0030\u0038\u0030","\u0063\u0030\u0063\u0030\u0063\u0030","\u0038\u0030\u0038\u0030\u0038\u0030","\u0039\u0039\u0039\u0039\u0066\u0066","\u0039\u0039\u0033\u0033\u0036\u0036","\u0066\u0066\u0066\u0066\u0063\u0063","\u0063\u0063\u0066\u0066\u0066\u0066","\u0036\u0036\u0030\u0030\u0036\u0036","\u0066\u0066\u0038\u0030\u0038\u0030","\u0030\u0030\u0036\u0036\u0063\u0063","\u0063\u0063\u0063\u0063\u0066\u0066","\u0030\u0030\u0030\u0030\u0038\u0030","\u0066\u0066\u0030\u0030\u0066\u0066","\u0066\u0066\u0066\u0066\u0030\u0030","\u0030\u0030\u0066\u0066\u0066\u0066","\u0038\u0030\u0030\u0030\u0038\u0030","\u0038\u0030\u0030\u0030\u0030\u0030","\u0030\u0030\u0038\u0030\u0038\u0030","\u0030\u0030\u0030\u0030\u0066\u0066","\u0030\u0030\u0063\u0063\u0066\u0066","\u0063\u0063\u0066\u0066\u0066\u0066","\u0063\u0063\u0066\u0066\u0063\u0063","\u0066\u0066\u0066\u0066\u0039\u0039","\u0039\u0039\u0063\u0063\u0066\u0066","\u0066\u0066\u0039\u0039\u0063\u0063","\u0063\u0063\u0039\u0039\u0066\u0066","\u0066\u0066\u0063\u0063\u0039\u0039","\u0033\u0033\u0036\u0036\u0066\u0066","\u0033\u0033\u0063\u0063\u0063\u0063","\u0039\u0039\u0063\u0063\u0030\u0030","\u0066\u0066\u0063\u0063\u0030\u0030","\u0066\u0066\u0039\u0039\u0030\u0030","\u0066\u0066\u0036\u0036\u0030\u0030","\u0036\u0036\u0036\u0036\u0039\u0039","\u0039\u0036\u0039\u0036\u0039\u0036","\u0030\u0030\u0033\u0033\u0036\u0036","\u0033\u0033\u0039\u0039\u0036\u0036","\u0030\u0030\u0033\u0033\u0030\u0030","\u0033\u0033\u0033\u0033\u0030\u0030","\u0039\u0039\u0033\u0033\u0030\u0030","\u0039\u0039\u0033\u0033\u0036\u0036","\u0033\u0033\u0033\u0033\u0039\u0039","\u0033\u0033\u0033\u0033\u0033\u0033"};

// RegisterFont makes a PdfFont accessible for using in converting to PDF.
func RegisterFont (name string ,style FontStyle ,font *_bb .PdfFont ){_ba .RegisterFont (name ,style ,font );};const (FontStyle_Regular FontStyle =0;FontStyle_Bold FontStyle =1;FontStyle_Italic FontStyle =2;FontStyle_BoldItalic FontStyle =3;);func (_cccg *convertContext )alignSymbolsHorizontally (_dcc *cell ,_fea _bc .ST_HorizontalAlignment ){if _fea ==_bc .ST_HorizontalAlignmentUnset {switch _dcc ._gab {case _bc .ST_CellTypeB :_fea =_bc .ST_HorizontalAlignmentCenter ;case _bc .ST_CellTypeN :_fea =_bc .ST_HorizontalAlignmentRight ;default:_fea =_bc .ST_HorizontalAlignmentLeft ;};};var _fgd float64 ;for _ ,_bda :=range _dcc ._ffce {switch _fea {case _bc .ST_HorizontalAlignmentLeft :_fgd =_de ;case _bc .ST_HorizontalAlignmentRight :_gfag :=_edb (_bda ._cbf );_fgd =_dcc ._bacg -_de -_gfag ;case _bc .ST_HorizontalAlignmentCenter :_gfee :=_edb (_bda ._cbf );_fgd =(_dcc ._bacg -_gfee )/2;};for _ ,_ggd :=range _bda ._cbf {_ggd ._caee +=_fgd ;};};};type colInfo struct{_gccb float64 ;_bec float64 ;_bdaa *style ;};func (_ccaa *convertContext )getStyle (_abcgd *uint32 )*style {_bgc :=&style {};_dbeg :=false ;if _abcgd !=nil {_aeb :=_ccaa ._bfgd .GetCellStyle (*_abcgd );_ggba :=_aeb .GetFont ();for _ ,_caea :=range _ggba .Name {if _caea !=nil {_bgc ._eade =&_caea .ValAttr ;_dbeg =true ;break ;};};for _ ,_bebe :=range _ggba .B {if _bebe !=nil {_cabf :=_bebe .ValAttr ==nil ||*_bebe .ValAttr ;_bgc ._cedd =&_cabf ;_dbeg =true ;break ;};};for _ ,_cbbb :=range _ggba .I {if _cbbb !=nil {_dgfdc :=_cbbb .ValAttr ==nil ||*_cbbb .ValAttr ;_bgc ._ecdf =&_dgfdc ;_dbeg =true ;break ;};};for _ ,_gca :=range _ggba .U {if _gca !=nil {_edfe :=_gca .ValAttr ==_bc .ST_UnderlineValuesSingle ||_gca .ValAttr ==_bc .ST_UnderlineValuesUnset ;_bgc ._ggbf =&_edfe ;_dbeg =true ;break ;};};for _ ,_bgf :=range _ggba .Sz {if _bgf !=nil {_ecf :=_bgf .ValAttr /12*_ba .DefaultFontSize ;_bgc ._gaedc =&_ecf ;_dbeg =true ;break ;};};for _ ,_dgff :=range _ggba .VertAlign {if _dgff !=nil {_cdcb :=_dgff .ValAttr ==_dd .ST_VerticalAlignRunSuperscript ;_bgc ._cfff =&_cdcb ;_cafc :=_dgff .ValAttr ==_dd .ST_VerticalAlignRunSubscript ;_bgc ._faa =&_cafc ;_dbeg =true ;break ;};};for _ ,_bfe :=range _ggba .Color {if _bfe !=nil {_bgc ._adfb =_ccaa .getColorStringFromSmlColor (_bfe );_dbeg =true ;break ;};};_baba :=_aeb .GetBorder ();if _baba .Top !=nil {_bgc ._bfbb =_ccaa .getBorder (_baba .Top );_dbeg =true ;};if _baba .Bottom !=nil {_bgc ._gece =_ccaa .getBorder (_baba .Bottom );_dbeg =true ;};if _baba .Left !=nil {_bgc ._gdfe =_ccaa .getBorder (_baba .Left );_dbeg =true ;};if _baba .Right !=nil {_bgc ._bag =_ccaa .getBorder (_baba .Right );_dbeg =true ;};if _aeb .Wrapped (){_bgc ._afeb =true ;_dbeg =true ;};if _bcd :=_aeb .GetVerticalAlignment ();_bcd !=_bc .ST_VerticalAlignmentUnset {_bgc ._geg =_bcd ;_dbeg =true ;};if _fcbd :=_aeb .GetHorizontalAlignment ();_fcbd !=_bc .ST_HorizontalAlignmentUnset {_bgc ._bebc =_fcbd ;_dbeg =true ;};};if _dbeg {return _bgc ;};return nil ;};var _fbc =_cac (1);

// RegisterFontsFromDirectory registers all fonts from the given directory automatically detecting font families and styles.
func RegisterFontsFromDirectory (dirName string )error {return _ba .RegisterFontsFromDirectory (dirName )};func (_ccaf *convertContext )combineCellStyleWithRPrElt (_fgab *style ,_afb *_bc .CT_RPrElt )*style {_aead :=*_fgab ;_afed :=_ccaf .getStyleFromRPrElt (_afb );if _afed ==nil {return &_aead ;};if _afed ._adfb !=nil {_aead ._adfb =_afed ._adfb ;};if _afed ._gaedc !=nil {_aead ._gaedc =_afed ._gaedc ;};if _afed ._eade !=nil {_aead ._eade =_afed ._eade ;};if _afed ._cedd !=nil {_aead ._cedd =_afed ._cedd ;};if _afed ._ecdf !=nil {_aead ._ecdf =_afed ._ecdf ;};if _afed ._ggbf !=nil {_aead ._ggbf =_afed ._ggbf ;};if _afed ._cfff !=nil {_aead ._cfff =_afed ._cfff ;};if _afed ._faa !=nil {_aead ._faa =_afed ._faa ;};return &_aead ;};func (_dab *convertContext )getColorFromTheme (_cfde uint32 )string {_cece :=_dab ._egag .Themes ();if len (_cece )!=0{_cbcc :=_cece [0];if _gdde :=_cbcc .ThemeElements ;_gdde !=nil {if _dcae :=_gdde .ClrScheme ;_dcae !=nil {switch _cfde {case 0:return _ba .GetColorStringFromDmlColor (_dcae .Lt1 );case 1:return _ba .GetColorStringFromDmlColor (_dcae .Dk1 );case 2:return _ba .GetColorStringFromDmlColor (_dcae .Lt2 );case 3:return _ba .GetColorStringFromDmlColor (_dcae .Dk2 );case 4:return _ba .GetColorStringFromDmlColor (_dcae .Accent1 );case 5:return _ba .GetColorStringFromDmlColor (_dcae .Accent2 );case 6:return _ba .GetColorStringFromDmlColor (_dcae .Accent3 );case 7:return _ba .GetColorStringFromDmlColor (_dcae .Accent4 );case 8:return _ba .GetColorStringFromDmlColor (_dcae .Accent5 );case 9:return _ba .GetColorStringFromDmlColor (_dcae .Accent6 );};};};};return "";};type symbol struct{_eabb string ;_caee float64 ;_agga float64 ;_aeab float64 ;_gee *_df .TextStyle ;_edab string ;};func (_aebd *convertContext )getBorder (_gbga *_bc .CT_BorderPr )*border {_bfcdb :=&border {};switch _gbga .StyleAttr {case _bc .ST_BorderStyleThin :_bfcdb ._eeg =_gb ;case _bc .ST_BorderStyleMedium :_bfcdb ._eeg =_gb *2;case _bc .ST_BorderStyleThick :_bfcdb ._eeg =_gb *4;};if _bfcdb ._eeg ==0.0{return nil ;};if _eabbd :=_gbga .Color ;_eabbd !=nil {_fad :=_aebd .getColorStringFromSmlColor (_eabbd );if _fad !=nil {_bfcdb ._bgd =_df .ColorRGBFromHex (*_fad );}else {_bfcdb ._bgd =_df .ColorBlack ;};};return _bfcdb ;};var _afebg =map[uint32 ]_df .PageSize {1:_df .PageSize {8.5*_e .Inch ,11*_e .Inch },2:_df .PageSize {8.5*_e .Inch ,11*_e .Inch },3:_df .PageSize {11*_e .Inch ,17*_e .Inch },4:_df .PageSize {17*_e .Inch ,11*_e .Inch },5:_df .PageSize {8.5*_e .Inch ,14*_e .Inch },6:_df .PageSize {5.5*_e .Inch ,8.5*_e .Inch },7:_df .PageSize {7.5*_e .Inch ,10*_e .Inch },8:_df .PageSize {_cac (297),_cac (420)},9:_df .PageSize {_cac (210),_cac (297)},10:_df .PageSize {_cac (210),_cac (297)},11:_df .PageSize {_cac (148),_cac (210)},70:_df .PageSize {_cac (105),_cac (148)},12:_df .PageSize {_cac (250),_cac (354)},13:_df .PageSize {_cac (182),_cac (257)},14:_df .PageSize {8.5*_e .Inch ,13*_e .Inch },20:_df .PageSize {4.125*_e .Inch ,9.5*_e .Inch },27:_df .PageSize {_cac (110),_cac (220)},28:_df .PageSize {_cac (162),_cac (229)},34:_df .PageSize {_cac (250),_cac (176)},29:_df .PageSize {_cac (324),_cac (458)},30:_df .PageSize {_cac (229),_cac (324)},31:_df .PageSize {_cac (114),_cac (162)},37:_df .PageSize {3.88*_e .Inch ,7.5*_e .Inch },43:_df .PageSize {_cac (100),_cac (148)},69:_df .PageSize {_cac (200),_cac (148)}};func _cac (_gcfg float64 )float64 {return _gcfg *_e .Millimeter };func (_agg *convertContext )makeMergedCells (){_cedc :=[]*mergedCell {};for _ ,_bcce :=range _agg ._ccbd .MergedCells (){_bbcb ,_agc ,_cfb :=_ag .ParseRangeReference (_bcce .Reference ());if _cfb !=nil {_c .Log .Debug ("\u0065\u0072r\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u006d\u0065\u0072\u0067\u0065\u0064\u0020\u0063\u0065\u006c\u006c: \u0025\u0073",_cfb );continue ;};_bba :=mergedCell {_aecb :_bbcb .RowIdx ,_becb :_bbcb .ColumnIdx ,_aee :_agc .RowIdx ,_bfcg :_agc .ColumnIdx };for _aaa :=_bba ._becb -1;_aaa < _bba ._bfcg ;_aaa ++{_bba ._gfeb +=_agg ._dgbc [_aaa ]._bec ;};for _adf :=_bba ._aecb -1;_adf < _bba ._aee ;_adf ++{_bba ._ebdec +=_agg ._gag [_adf ]._ccdb ;};_cedc =append (_cedc ,&_bba );};_agg ._eacd =_cedc ;};func (_edc *convertContext )getContentFromCell (_bac _cd .Cell ,_gbcf *style ,_fbg float64 ,_cceg bool )([]*line ,_bc .ST_CellType ){_dege :=_bac .X ();var _bbcbf []*symbol ;switch _dege .TAttr {case _bc .ST_CellTypeS :_dda :=_dege .V ;if _dda !=nil {_dbag ,_bfc :=_d .Atoi (*_dda );if _bfc ==nil {_ddae :=_edc ._egag .SharedStrings .X ().Si [_dbag ];if _ddae .T !=nil {_bbcbf =_edc .getSymbolsFromString (*_ddae .T ,_gbcf );}else if _ddae .R !=nil {_bbcbf =_edc .getSymbolsFromR (_ddae .R ,_gbcf );};};};case _bc .ST_CellTypeB :_dbg :=_dege .V ;if _dbg !=nil {if *_dbg =="\u0030"{_bbcbf =_edc .getSymbolsFromString ("\u0046\u0041\u004cS\u0045",_gbcf );}else {_bbcbf =_edc .getSymbolsFromString ("\u0054\u0052\u0055\u0045",_gbcf );};};default:_bbcbf =_edc .getSymbolsFromString (_bac .GetFormattedValue (),_gbcf );};_beba :=0.0;_aba :=0.0;var _eaf []*line ;var _afc bool ;if _gbcf !=nil {if _gbcf ._cfff !=nil {if *_gbcf ._cfff {_afc =true ;};};if _gbcf ._faa !=nil {if *_gbcf ._faa {_afc =true ;};};};if _cceg {_eaf =[]*line {};_afeg :=_fbg -2*_de ;_bde :=[]*symbol {};for _ ,_gggc :=range _bbcbf {_adb (_gggc );if _beba +_gggc ._aeab >=_afeg {_adaae :=_daa (_bde );if _afc {_adaae /=_f ;};_eaf =append (_eaf ,&line {_edfb :_aba ,_cbf :_bde ,_ffgf :_adaae });_bde =[]*symbol {_gggc };_beba =_gggc ._aeab ;_aba +=_adaae ;}else {_gggc ._caee =_beba ;_beba +=_gggc ._aeab ;_bde =append (_bde ,_gggc );};};_cagc :=_daa (_bde );if _afc {_cagc /=_f ;};if len (_bde )> 0{_eaf =append (_eaf ,&line {_edfb :_aba ,_cbf :_bde ,_ffgf :_cagc });};}else {for _ ,_bgea :=range _bbcbf {_adb (_bgea );_bgea ._caee =_beba ;_beba +=_bgea ._aeab ;};if len (_bbcbf )> 0{_eaf =[]*line {&line {_cbf :_bbcbf ,_ffgf :_daa (_bbcbf )}};};};_gaed :=_dege .TAttr ;if _gaed ==_bc .ST_CellTypeUnset {_gaed =_bc .ST_CellTypeN ;};return _eaf ,_gaed ;};func _edb (_ecg []*symbol )float64 {_dbga :=0.0;for _ ,_acd :=range _ecg {_dbga +=_acd ._aeab ;};return _dbga ;};type convertContext struct{_cbaf *_df .Creator ;_egag *_cd .Workbook ;_ffb *_gd .Theme ;_ccbd *_cd .Sheet ;_bfgd *_cd .StyleSheet ;_ddeg int ;_fgdb int ;_gffe []*pagespan ;_dgbe *page ;_dgbc []*colInfo ;_gag []*rowInfo ;_cegg []*rowspan ;_acdc float64 ;_efae float64 ;_abaa float64 ;_dgae float64 ;_eacd []*mergedCell ;_bdb []*anchor ;};func (_agcf *convertContext )getSymbolsFromString (_fff string ,_acc *style )[]*symbol {_ebfc :=[]*symbol {};_bcg :=_agcf .makeTextStyleFromCellStyle (_acc );for _ ,_ffef :=range _fff {_ebfc =append (_ebfc ,&symbol {_eabb :string (_ffef ),_gee :_bcg });};return _ebfc ;};func (_adcb *convertContext )getSymbolsFromR (_bca []*_bc .CT_RElt ,_bab *style )[]*symbol {_afg :=[]*symbol {};for _ ,_eacg :=range _bca {_edec :=_adcb .combineCellStyleWithRPrElt (_bab ,_eacg .RPr );for _ ,_agaa :=range _eacg .T {_afg =append (_afg ,&symbol {_eabb :string (_agaa ),_gee :_adcb .makeTextStyleFromCellStyle (_edec )});};};return _afg ;};var _gb =_cac (0.0625);const _f =0.64;func (_ceb *convertContext )alignSymbolsVertically (_bafd *cell ,_gfae _bc .ST_VerticalAlignment ){var _fee float64 ;switch _gfae {case _bc .ST_VerticalAlignmentTop :_fee =_fb ;if _bafd ._bagc {_fee -=_eb ;}else if _bafd ._dbe {_fee +=4*_eb ;};for _ ,_aaff :=range _bafd ._ffce {_fee +=_aaff ._ffgf ;_aaff ._edfb =_fee ;_fee +=_fbc ;};case _bc .ST_VerticalAlignmentCenter :_dedab :=0.0;for _ ,_cagb :=range _bafd ._ffce {_dedab +=_cagb ._ffgf +_cac (1);};_fee =0.5*(_bafd ._daea -_dedab );if _bafd ._bagc {_fee -=2*_eb ;}else if _bafd ._dbe {_fee +=2*_eb ;};for _ ,_ggb :=range _bafd ._ffce {_fee +=_ggb ._ffgf +0.5*_fbc ;_ggb ._edfb =_fee ;_fee +=0.5*_fbc ;};default:_fee =_bafd ._daea -_fb ;if _bafd ._bagc {_fee -=4*_eb ;}else if _bafd ._dbe {_fee +=_eb ;};for _fbd :=len (_bafd ._ffce )-1;_fbd >=0;_fbd --{_bafd ._ffce [_fbd ]._edfb =_fee ;_fee -=_bafd ._ffce [_fbd ]._ffgf ;_fee -=_fbc ;};};};func (_gbe *convertContext )makeRows (){_aa :=[]*rowInfo {};_eff :=_gbe ._ccbd .Rows ();_gfe :=0;for _ ,_gcf :=range _eff {_gfe ++;_cba :=int (_gcf .RowNumber ());if _cba > _gfe {for _bcc :=_gfe ;_bcc < _cba ;_bcc ++{_aa =append (_aa ,&rowInfo {_ccdb :16/_ebf });};_gfe =_cba ;};var _bdg float64 ;if _gcf .X ().HtAttr ==nil {_bdg =16;}else {_bdg =*_gcf .X ().HtAttr ;};_aa =append (_aa ,&rowInfo {_ccdb :_bdg /_ebf ,_bfdc :true ,_adaaf :_gbe .getStyle (_gcf .X ().SAttr )});};for _ded :=len (_aa );_ded < _gbe ._ddeg ;_ded ++{_aa =append (_aa ,&rowInfo {_ccdb :16/_ebf });};_gbe ._gag =_aa ;};type style struct{_adfb *string ;_gaedc *float64 ;_eade *string ;_cedd *bool ;_ecdf *bool ;_ggbf *bool ;_cfff *bool ;_faa *bool ;_bfbb *border ;_gece *border ;_gdfe *border ;_bag *border ;_afeb bool ;_geg _bc .ST_VerticalAlignment ;_bebc _bc .ST_HorizontalAlignment ;};func (_dfdf *convertContext )getStyleFromCell (_eceb _cd .Cell ,_eeee ,_efba *style )*style {_bebcg :=_eceb .X ();_ddca :=_dfdf .getStyle (_bebcg .SAttr );_ebdf (_ddca ,_eeee );_ebdf (_ddca ,_efba );return _ddca ;};func (_bbc *convertContext )determineMaxIndexes (){var _fgb ,_ee int ;_fgb =int (_bbc ._ccbd .MaxColumnIdx ());_agbe :=_bbc ._ccbd .Rows ();if len (_agbe )> 0{_ee =int (_agbe [len (_agbe )-1].RowNumber ());};for _ ,_fcf :=range _bbc ._bdb {if _fcf ._cgf >=_ee {_ee =_fcf ._cgf +1;};if _fcf ._dbgc >=_fgb {_fgb =_fcf ._dbgc +1;};};_bbc ._ddeg =_ee ;_bbc ._fgdb =_fgb ;};func (_eaff *convertContext )makeTextStyleFromCellStyle (_effcb *style )*_df .TextStyle {_aeeg :=_eaff ._cbaf .NewTextStyle ();if _effcb ==nil {_aeeg .FontSize =_ba .DefaultFontSize ;_aeeg .Font =_ba .AssignStdFontByName (_aeeg ,_ba .StdFontsMap ["\u0064e\u0066\u0061\u0075\u006c\u0074"][FontStyle_Regular ]);return &_aeeg ;};if _fecd (_effcb ._ggbf ){_aeeg .Underline =true ;_aeeg .UnderlineStyle =_df .TextDecorationLineStyle {Offset :0.5,Thickness :_cac (1/32)};};var _ceae FontStyle ;if _fecd (_effcb ._cedd )&&_fecd (_effcb ._ecdf ){_ceae =FontStyle_BoldItalic ;}else if _fecd (_effcb ._cedd ){_ceae =FontStyle_Bold ;}else if _fecd (_effcb ._ecdf ){_ceae =FontStyle_Italic ;}else {_ceae =FontStyle_Regular ;};_deaa :="\u0064e\u0066\u0061\u0075\u006c\u0074";if _effcb ._eade !=nil {_deaa =*_effcb ._eade ;};if _ebad ,_bcb :=_ba .StdFontsMap [_deaa ];_bcb {_aeeg .Font =_ba .AssignStdFontByName (_aeeg ,_ebad [_ceae ]);}else if _efeg :=_ba .GetRegisteredFont (_deaa ,_ceae );_efeg !=nil {_aeeg .Font =_efeg ;}else {_c .Log .Debug ("\u0046\u006f\u006e\u0074\u0020\u0025\u0073\u0020\u0077\u0069\u0074h\u0020\u0073\u0074\u0079\u006c\u0065\u0020\u0025s\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002c\u0020\u0072\u0065\u0073\u0065\u0074 \u0074\u006f\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u002e",_deaa ,_ceae );_aeeg .Font =_ba .AssignStdFontByName (_aeeg ,_ba .StdFontsMap ["\u0064e\u0066\u0061\u0075\u006c\u0074"][_ceae ]);};if _effcb ._gaedc !=nil {_aeeg .FontSize =*_effcb ._gaedc ;};if _effcb ._adfb !=nil {_aeeg .Color =_df .ColorRGBFromHex (*_effcb ._adfb );};if _effcb ._cfff !=nil &&*_effcb ._cfff {_aeeg .FontSize *=_f ;}else if _effcb ._faa !=nil &&*_effcb ._faa {_aeeg .FontSize *=_f ;};return &_aeeg ;};

// ConvertToPdf converts a sheet to a PDF file. This package is beta, breaking changes can take place.
func ConvertToPdf (s *_cd .Sheet )*_df .Creator {_ac :=s .X ();if _ac ==nil {return nil ;};var _fa _df .PageSize ;var _fg bool ;if _dg :=_ac .PageSetup ;_dg !=nil {if _ge :=_dg .PaperSizeAttr ;_ge !=nil {_fa =_afebg [*_ge ];};_fg =_dg .OrientationAttr ==_bc .ST_OrientationLandscape ;};if (_fa ==_df .PageSize {}){_fa =_afebg [1];};if _fg {_fa [0],_fa [1]=_fa [1],_fa [0];};_gc :=_df .New ();_gc .SetPageSize (_fa );var _ae ,_ec ,_ed ,_gg float64 ;if _cde :=_ac .PageMargins ;_cde !=nil {_ed =_cde .LeftAttr ;_gg =_cde .RightAttr ;_ae =_cde .TopAttr ;_ec =_cde .BottomAttr ;};if _ed < 0.25{_ed =0.25;};if _gg < 0.25{_gg =0.25;};_ae *=_e .Inch ;_ec *=_e .Inch ;_ed *=_e .Inch ;_gg *=_e .Inch ;_gc .SetPageMargins (_ed ,_gg ,_ae ,_ec );_dfa :=s .Workbook ();var _ged *_gd .Theme ;if len (_dfa .Themes ())> 0{_ged =_dfa .Themes ()[0];};_ad :=&convertContext {_cbaf :_gc ,_ccbd :s ,_egag :s .Workbook (),_ffb :_ged ,_bfgd :&s .Workbook ().StyleSheet ,_acdc :_ae ,_efae :_ed ,_abaa :_fa [1]-_ec -_ae ,_dgae :_fa [0]-_gg -_ed };_ad .makeAnchors ();_ad .determineMaxIndexes ();if _ad ._ddeg ==0&&_ad ._fgdb ==0{_gc .NewPage ();return _gc ;};_ad .makeCols ();_ad .makeRows ();_ad .makeMergedCells ();_ad .makeCells ();_ad .makePagespans ();_ad .makeRowspans ();_ad .makePages ();_ad .fillPages ();_ad .distributeAnchors ();_ad .drawSheet ();return _gc ;};func _daa (_gdfa []*symbol )float64 {_dcb :=0.0;for _ ,_daaa :=range _gdfa {if _daaa ._agga > _dcb {_dcb =_daaa ._agga ;};};return _dcb ;};func (_bdgg *convertContext )addRowToPage (_adaa []*cell ,_fgbd int ){_efb :=0.0;_gae :=_bdgg ._dgae ;for _ ,_gfca :=range _adaa {if len (_gfca ._ffce )!=0{_gfca ._gdag =_efb ;_efb =_gfca ._accd +_gfca ._bacg ;};};for _gcfe :=len (_adaa )-1;_gcfe >=0;_gcfe --{_ebca :=_adaa [_gcfe ];if len (_ebca ._ffce )!=0{_ebca ._agf =_gae ;_gae =_ebca ._accd ;};};_bdgg ._dgbe ._dbaa =append (_bdgg ._dgbe ._dbaa ,&pageRow {_cbd :_fgbd ,_adacf :_adaa });};func _fecd (_eeef *bool )bool {return _eeef !=nil &&*_eeef };const _de =3;type line struct{_edfb float64 ;_cbf []*symbol ;_ffgf float64 ;};type anchor struct{_gbgd _g .Image ;_dggb *_a .ChartSpace ;_acga int ;_agdg int64 ;_gdd int ;_caba int64 ;_cgf int ;_fd int64 ;_dbgc int ;_fbcb int64 ;};func (_eg *convertContext )makeAnchors (){_agd ,_ece :=_eg ._ccbd .GetDrawing ();if _agd !=nil {for _ ,_bd :=range _agd .EG_Anchor {_ced :=&anchor {};if _bg :=_bd .TwoCellAnchor ;_bg !=nil {_gce ,_ddd :=_bg .From ,_bg .To ;if _gce ==nil ||_ddd ==nil {return ;};_ced ._acga =int (_gce .Row );_ced ._agdg =_ba .FromSTCoordinate (_gce .RowOff );_ced ._gdd =int (_gce .Col );_ced ._caba =_ba .FromSTCoordinate (_gce .ColOff );_ced ._cgf =int (_ddd .Row );_ced ._fd =_ba .FromSTCoordinate (_ddd .RowOff );_ced ._dbgc =int (_ddd .Col );_ced ._fbcb =_ba .FromSTCoordinate (_ddd .ColOff );if _bge :=_bg .Choice ;_bge !=nil {if _bgb :=_bge .Pic ;_bgb !=nil {if _eba :=_bgb .BlipFill ;_eba !=nil {if _bf :=_eba .Blip ;_bf !=nil {if _bdc :=_bf .EmbedAttr ;_bdc !=nil {for _ ,_eca :=range _ece .X ().Relationship {if _eca .IdAttr ==*_bdc {for _ ,_acg :=range _eg ._egag .Images {if _acg .Target ()==_eca .TargetAttr {_ddf ,_gbb :=_ef .Open (_acg .Path ());if _gbb !=nil {_c .Log .Debug ("\u004fp\u0065\u006e\u0020\u0069m\u0061\u0067\u0065\u0020\u0066i\u006ce\u0020e\u0072\u0072\u006f\u0072\u003a\u0020\u0025s",_gbb );continue ;};_gceg ,_ ,_gbb :=_g .Decode (_ddf );if _gbb !=nil {_c .Log .Debug ("\u0044\u0065\u0063\u006fde\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020%\u0073",_gbb );continue ;};_ced ._gbgd =_gceg ;};};};};};};};}else if _bdd :=_bge .GraphicFrame ;_bdd !=nil {if _dddb :=_bdd .Graphic ;_dddb !=nil {if _ecc :=_dddb .GraphicData ;_ecc !=nil {for _ ,_cdg :=range _ecc .Any {if _ada ,_acf :=_cdg .(*_a .Chart );_acf {for _ ,_cc :=range _ece .X ().Relationship {if _cc .IdAttr ==_ada .IdAttr {_ga :=_eg ._egag .GetChartByTargetId (_cc .TargetAttr );if _ga !=nil {_ced ._dggb =_ga ;};};};};};};};};};};if _ced ._gbgd !=nil ||_ced ._dggb !=nil {_eg ._bdb =append (_eg ._bdb ,_ced );};};};};func (_gbg *convertContext )makePagespans (){_gbg ._gffe =[]*pagespan {};_bed :=0.0;_deg :=0;for _eef ,_gcg :=range _gbg ._dgbc {_eee :=_gcg ._bec ;if _bed +_eee <=_gbg ._dgae {_gcg ._gccb =_bed ;_bed +=_eee ;}else {_gcg ._gccb =0;_gbg ._gffe =append (_gbg ._gffe ,&pagespan {_aggb :_bed ,_gead :_deg ,_bgbad :_eef });_bed =_eee ;_deg =_eef ;};};_gbg ._gffe =append (_gbg ._gffe ,&pagespan {_aggb :_bed ,_gead :_deg ,_bgbad :len (_gbg ._dgbc )});};type border struct{_eeg float64 ;_bgd _df .Color ;};func (_dadg *convertContext )drawPage (_ffga *page ){_cag :=_dadg ._acdc ;_ede :=_dadg ._efae ;for _ ,_ffgb :=range _ffga ._dbaa {_dag :=_dadg ._gag [_ffgb ._cbd ];for _ ,_fge :=range _ffgb ._adacf {_bbbd :=_fge ._gdag < _fge ._accd ;_fgg :=_fge ._agf > _fge ._accd +_fge ._bacg ;var _cgge ,_fgc bool ;for _ ,_bdcf :=range _fge ._ffce {for _ ,_ddb :=range _bdcf ._cbf {if _bbbd &&!_cgge {_cgge =_ddb ._caee < 0;};if _fgg &&!_fgc {_fgc =_fge ._bacg < _ddb ._caee +_ddb ._aeab ;};if _fge ._accd +_ddb ._caee >=_fge ._gdag &&_fge ._accd +_ddb ._caee +_ddb ._aeab <=_fge ._agf {_ead :=_dadg ._cbaf .NewStyledParagraph ();_gda :=_ede +_fge ._accd +_ddb ._caee ;_dbfe :=_cag +_dag ._ecde +_bdcf ._edfb -_ddb ._agga -_cac (0.5);_ead .SetPos (_gda ,_dbfe );var _gggd *_df .TextChunk ;if _ddb ._edab !=""{_gggd =_ead .AddExternalLink (_ddb ._eabb ,_ddb ._edab );}else {_gggd =_ead .Append (_ddb ._eabb );};if _ddb ._gee !=nil {_gggd .Style =*_ddb ._gee ;};_dadg ._cbaf .Draw (_ead );};};};var _aef ,_cdea ,_agaf ,_egd ,_abcg ,_efe float64 ;var _cdb ,_cfe ,_bebd ,_dcg _df .Color ;if _ebg :=_fge ._abe ;_ebg !=nil {_aef =_ebg ._eeg ;_cdb =_ebg ._bgd ;};if _eda :=_fge ._fged ;_eda !=nil {_cdea =_eda ._eeg ;_cfe =_eda ._bgd ;};if _egaa :=_fge ._dgg ;_egaa !=nil {_agaf =_egaa ._eeg ;_abcg =_agaf /2;_bebd =_egaa ._bgd ;};if _cgd :=_fge ._bfcd ;_cgd !=nil {_egd =_cgd ._eeg ;_efe =_egd /2;_dcg =_cgd ._bgd ;};var _ecag float64 ;if _ffgb ._cbd > 1{_ecag =_dadg ._gag [_ffgb ._cbd -1]._fec ;};_fcce :=_cag +_dag ._ecde -0.5*(_ecag -_aef );_gbfc :=_cag +_dag ._ecde +_dag ._ccdb +0.5*(_dag ._fec +_cdea );_gdab :=_ede +_fge ._accd ;_fgf :=_gdab +_fge ._cbad ;_ba .DrawLine (_dadg ._cbaf ,_gdab ,_fcce ,_fgf ,_fcce ,_aef ,_cdb );_ba .DrawLine (_dadg ._cbaf ,_gdab ,_gbfc ,_fgf ,_gbfc ,_cdea ,_cfe );if !_cgge {_ba .DrawLine (_dadg ._cbaf ,_gdab -_abcg ,_fcce ,_gdab -_abcg ,_gbfc ,_agaf ,_bebd );};if !_fgc {_ba .DrawLine (_dadg ._cbaf ,_fgf -_efe ,_fcce ,_fgf -_efe ,_gbfc ,_egd ,_dcg );};};};for _ ,_ebe :=range _ffga ._abbg {if _ebe !=nil {_dadg ._cbaf .Draw (_ebe );};};};type rowspan struct{_ebcd float64 ;_ccagg int ;_geb int ;};type pageRow struct{_cbd int ;_adacf []*cell ;};