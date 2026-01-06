package LogicExamples

import "gocurry"
import "curry2go/Algorithms"
import "curry2go/Network"
import "curry2go/Prelude"

var func_names []string = []string{ "lEARNING_RATE", "andExamples", "andLayers", "andWeights", "andNetwork", "orExamples", "orLayers", "orWeights", "orNetwork", "notExamples", "notLayers", "notWeights", "notNetwork", "nandExamples", "nandLayers", "nandWeights", "nandNetwork", "xorExamples", "xorLayers", "xorWeights", "xorNetwork", "xorplusExamples", "xorplusLayers", "xorplusWeights", "xorplusNetwork" }

func LogicExamples__CREATE_lEARNINGUs_RATE( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, LogicExamples_lEARNINGUs_RATE, & func_names[ 0 ], 0, -1, args... )
    return( root )
}

func LogicExamples__CREATE_andExamples( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, LogicExamples_andExamples, & func_names[ 1 ], 0, -1, args... )
    return( root )
}

func LogicExamples__CREATE_andLayers( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, LogicExamples_andLayers, & func_names[ 2 ], 0, -1, args... )
    return( root )
}

func LogicExamples__CREATE_andWeights( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, LogicExamples_andWeights, & func_names[ 3 ], 0, -1, args... )
    return( root )
}

func LogicExamples__CREATE_andNetwork( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, LogicExamples_andNetwork, & func_names[ 4 ], 0, -1, args... )
    return( root )
}

func LogicExamples__CREATE_orExamples( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, LogicExamples_orExamples, & func_names[ 5 ], 0, -1, args... )
    return( root )
}

func LogicExamples__CREATE_orLayers( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, LogicExamples_orLayers, & func_names[ 6 ], 0, -1, args... )
    return( root )
}

func LogicExamples__CREATE_orWeights( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, LogicExamples_orWeights, & func_names[ 7 ], 0, -1, args... )
    return( root )
}

func LogicExamples__CREATE_orNetwork( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, LogicExamples_orNetwork, & func_names[ 8 ], 0, -1, args... )
    return( root )
}

func LogicExamples__CREATE_notExamples( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, LogicExamples_notExamples, & func_names[ 9 ], 0, -1, args... )
    return( root )
}

func LogicExamples__CREATE_notLayers( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, LogicExamples_notLayers, & func_names[ 10 ], 0, -1, args... )
    return( root )
}

func LogicExamples__CREATE_notWeights( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, LogicExamples_notWeights, & func_names[ 11 ], 0, -1, args... )
    return( root )
}

func LogicExamples__CREATE_notNetwork( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, LogicExamples_notNetwork, & func_names[ 12 ], 0, -1, args... )
    return( root )
}

func LogicExamples__CREATE_nandExamples( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, LogicExamples_nandExamples, & func_names[ 13 ], 0, -1, args... )
    return( root )
}

func LogicExamples__CREATE_nandLayers( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, LogicExamples_nandLayers, & func_names[ 14 ], 0, -1, args... )
    return( root )
}

func LogicExamples__CREATE_nandWeights( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, LogicExamples_nandWeights, & func_names[ 15 ], 0, -1, args... )
    return( root )
}

func LogicExamples__CREATE_nandNetwork( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, LogicExamples_nandNetwork, & func_names[ 16 ], 0, -1, args... )
    return( root )
}

func LogicExamples__CREATE_xorExamples( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, LogicExamples_xorExamples, & func_names[ 17 ], 0, -1, args... )
    return( root )
}

func LogicExamples__CREATE_xorLayers( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, LogicExamples_xorLayers, & func_names[ 18 ], 0, -1, args... )
    return( root )
}

func LogicExamples__CREATE_xorWeights( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, LogicExamples_xorWeights, & func_names[ 19 ], 0, -1, args... )
    return( root )
}

func LogicExamples__CREATE_xorNetwork( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, LogicExamples_xorNetwork, & func_names[ 20 ], 0, -1, args... )
    return( root )
}

func LogicExamples__CREATE_xorplusExamples( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, LogicExamples_xorplusExamples, & func_names[ 21 ], 0, -1, args... )
    return( root )
}

func LogicExamples__CREATE_xorplusLayers( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, LogicExamples_xorplusLayers, & func_names[ 22 ], 0, -1, args... )
    return( root )
}

func LogicExamples__CREATE_xorplusWeights( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, LogicExamples_xorplusWeights, & func_names[ 23 ], 0, -1, args... )
    return( root )
}

func LogicExamples__CREATE_xorplusNetwork( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, LogicExamples_xorplusNetwork, & func_names[ 24 ], 0, -1, args... )
    return( root )
}

func LogicExamples_lEARNINGUs_RATE( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    gocurry.FloatLitCreate( root, 0.1 )
    return
}

func LogicExamples_andExamples( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Prelude.Prelude__CREATE_Col_( root, Network.Network__CREATE_Example( root.NewNode(  ), gocurry.StringCreate( root.NewNode(  ), gocurry.ParseString( "FF" ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Network.Network__CREATE_Example( root.NewNode(  ), gocurry.StringCreate( root.NewNode(  ), gocurry.ParseString( "FT" ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Network.Network__CREATE_Example( root.NewNode(  ), gocurry.StringCreate( root.NewNode(  ), gocurry.ParseString( "TF" ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Network.Network__CREATE_Example( root.NewNode(  ), gocurry.StringCreate( root.NewNode(  ), gocurry.ParseString( "TT" ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) ) )
    return
}

func LogicExamples_andLayers( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Prelude.Prelude__CREATE_Col_( root, Network.Network__CREATE_Layer( root.NewNode(  ), gocurry.IntLitCreate( root.NewNode(  ), 1 ), Algorithms.Algorithms__CREATE_classifier( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) )
    return
}

func LogicExamples_andWeights( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Prelude.Prelude__CREATE_Col_( root, Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_implHash_negateHash_PreludeDot_NumHash_PreludeDot_FloatHash_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1.5 ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) )
    return
}

func LogicExamples_andNetwork( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Network.Network__CREATE_Network( root, LogicExamples__CREATE_andWeights( root.NewNode(  ) ), LogicExamples__CREATE_andExamples( root.NewNode(  ) ), LogicExamples__CREATE_andLayers( root.NewNode(  ) ), LogicExamples__CREATE_lEARNINGUs_RATE( root.NewNode(  ) ) )
    return
}

func LogicExamples_orExamples( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Prelude.Prelude__CREATE_Col_( root, Network.Network__CREATE_Example( root.NewNode(  ), gocurry.StringCreate( root.NewNode(  ), gocurry.ParseString( "FF" ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Network.Network__CREATE_Example( root.NewNode(  ), gocurry.StringCreate( root.NewNode(  ), gocurry.ParseString( "FT" ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Network.Network__CREATE_Example( root.NewNode(  ), gocurry.StringCreate( root.NewNode(  ), gocurry.ParseString( "TF" ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Network.Network__CREATE_Example( root.NewNode(  ), gocurry.StringCreate( root.NewNode(  ), gocurry.ParseString( "TT" ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) ) )
    return
}

func LogicExamples_orLayers( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Prelude.Prelude__CREATE_Col_( root, Network.Network__CREATE_Layer( root.NewNode(  ), gocurry.IntLitCreate( root.NewNode(  ), 1 ), Algorithms.Algorithms__CREATE_classifier( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) )
    return
}

func LogicExamples_orWeights( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Prelude.Prelude__CREATE_Col_( root, Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_implHash_negateHash_PreludeDot_NumHash_PreludeDot_FloatHash_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0.5 ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) )
    return
}

func LogicExamples_orNetwork( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Network.Network__CREATE_Network( root, LogicExamples__CREATE_orWeights( root.NewNode(  ) ), LogicExamples__CREATE_orExamples( root.NewNode(  ) ), LogicExamples__CREATE_orLayers( root.NewNode(  ) ), LogicExamples__CREATE_lEARNINGUs_RATE( root.NewNode(  ) ) )
    return
}

func LogicExamples_notExamples( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Prelude.Prelude__CREATE_Col_( root, Network.Network__CREATE_Example( root.NewNode(  ), gocurry.StringCreate( root.NewNode(  ), gocurry.ParseString( "F" ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Network.Network__CREATE_Example( root.NewNode(  ), gocurry.StringCreate( root.NewNode(  ), gocurry.ParseString( "T" ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) )
    return
}

func LogicExamples_notLayers( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Prelude.Prelude__CREATE_Col_( root, Network.Network__CREATE_Layer( root.NewNode(  ), gocurry.IntLitCreate( root.NewNode(  ), 1 ), Algorithms.Algorithms__CREATE_classifier( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) )
    return
}

func LogicExamples_notWeights( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Prelude.Prelude__CREATE_Col_( root, Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_implHash_negateHash_PreludeDot_NumHash_PreludeDot_FloatHash_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0.5 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) )
    return
}

func LogicExamples_notNetwork( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Network.Network__CREATE_Network( root, LogicExamples__CREATE_notWeights( root.NewNode(  ) ), LogicExamples__CREATE_notExamples( root.NewNode(  ) ), LogicExamples__CREATE_notLayers( root.NewNode(  ) ), LogicExamples__CREATE_lEARNINGUs_RATE( root.NewNode(  ) ) )
    return
}

func LogicExamples_nandExamples( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Prelude.Prelude__CREATE_Col_( root, Network.Network__CREATE_Example( root.NewNode(  ), gocurry.StringCreate( root.NewNode(  ), gocurry.ParseString( "FF" ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Network.Network__CREATE_Example( root.NewNode(  ), gocurry.StringCreate( root.NewNode(  ), gocurry.ParseString( "FT" ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Network.Network__CREATE_Example( root.NewNode(  ), gocurry.StringCreate( root.NewNode(  ), gocurry.ParseString( "TF" ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Network.Network__CREATE_Example( root.NewNode(  ), gocurry.StringCreate( root.NewNode(  ), gocurry.ParseString( "TT" ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) ) )
    return
}

func LogicExamples_nandLayers( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Prelude.Prelude__CREATE_Col_( root, Network.Network__CREATE_Layer( root.NewNode(  ), gocurry.IntLitCreate( root.NewNode(  ), 1 ), Algorithms.Algorithms__CREATE_logistic( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Network.Network__CREATE_Layer( root.NewNode(  ), gocurry.IntLitCreate( root.NewNode(  ), 1 ), Algorithms.Algorithms__CREATE_classifier( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) )
    return
}

func LogicExamples_nandWeights( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Prelude.Prelude__CREATE_Col_( root, Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_implHash_negateHash_PreludeDot_NumHash_PreludeDot_FloatHash_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1.5 ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_implHash_negateHash_PreludeDot_NumHash_PreludeDot_FloatHash_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0.5 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) )
    return
}

func LogicExamples_nandNetwork( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Network.Network__CREATE_Network( root, LogicExamples__CREATE_nandWeights( root.NewNode(  ) ), LogicExamples__CREATE_nandExamples( root.NewNode(  ) ), LogicExamples__CREATE_nandLayers( root.NewNode(  ) ), LogicExamples__CREATE_lEARNINGUs_RATE( root.NewNode(  ) ) )
    return
}

func LogicExamples_xorExamples( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Prelude.Prelude__CREATE_Col_( root, Network.Network__CREATE_Example( root.NewNode(  ), gocurry.StringCreate( root.NewNode(  ), gocurry.ParseString( "FF" ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Network.Network__CREATE_Example( root.NewNode(  ), gocurry.StringCreate( root.NewNode(  ), gocurry.ParseString( "FT" ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Network.Network__CREATE_Example( root.NewNode(  ), gocurry.StringCreate( root.NewNode(  ), gocurry.ParseString( "TF" ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Network.Network__CREATE_Example( root.NewNode(  ), gocurry.StringCreate( root.NewNode(  ), gocurry.ParseString( "TT" ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) ) )
    return
}

func LogicExamples_xorLayers( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Prelude.Prelude__CREATE_Col_( root, Network.Network__CREATE_Layer( root.NewNode(  ), gocurry.IntLitCreate( root.NewNode(  ), 2 ), Algorithms.Algorithms__CREATE_logistic( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Network.Network__CREATE_Layer( root.NewNode(  ), gocurry.IntLitCreate( root.NewNode(  ), 1 ), Algorithms.Algorithms__CREATE_classifier( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) )
    return
}

func LogicExamples_xorWeights( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Prelude.Prelude__CREATE_Col_( root, Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) )
    return
}

func LogicExamples_xorNetwork( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Network.Network__CREATE_Network( root, LogicExamples__CREATE_xorWeights( root.NewNode(  ) ), LogicExamples__CREATE_xorExamples( root.NewNode(  ) ), LogicExamples__CREATE_xorLayers( root.NewNode(  ) ), LogicExamples__CREATE_lEARNINGUs_RATE( root.NewNode(  ) ) )
    return
}

func LogicExamples_xorplusExamples( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Prelude.Prelude__CREATE_Col_( root, Network.Network__CREATE_Example( root.NewNode(  ), gocurry.StringCreate( root.NewNode(  ), gocurry.ParseString( "FF" ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Network.Network__CREATE_Example( root.NewNode(  ), gocurry.StringCreate( root.NewNode(  ), gocurry.ParseString( "FT" ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Network.Network__CREATE_Example( root.NewNode(  ), gocurry.StringCreate( root.NewNode(  ), gocurry.ParseString( "TF" ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Network.Network__CREATE_Example( root.NewNode(  ), gocurry.StringCreate( root.NewNode(  ), gocurry.ParseString( "TT" ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) ) )
    return
}

func LogicExamples_xorplusLayers( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Prelude.Prelude__CREATE_Col_( root, Network.Network__CREATE_Layer( root.NewNode(  ), gocurry.IntLitCreate( root.NewNode(  ), 4 ), Algorithms.Algorithms__CREATE_logistic( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Network.Network__CREATE_Layer( root.NewNode(  ), gocurry.IntLitCreate( root.NewNode(  ), 3 ), Algorithms.Algorithms__CREATE_logistic( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Network.Network__CREATE_Layer( root.NewNode(  ), gocurry.IntLitCreate( root.NewNode(  ), 2 ), Algorithms.Algorithms__CREATE_classifier( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) )
    return
}

func LogicExamples_xorplusWeights( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Prelude.Prelude__CREATE_Col_( root, Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) )
    return
}

func LogicExamples_xorplusNetwork( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Network.Network__CREATE_Network( root, LogicExamples__CREATE_xorplusWeights( root.NewNode(  ) ), LogicExamples__CREATE_xorplusExamples( root.NewNode(  ) ), LogicExamples__CREATE_xorplusLayers( root.NewNode(  ) ), LogicExamples__CREATE_lEARNINGUs_RATE( root.NewNode(  ) ) )
    return
}

