package TestPerceptron

import "gocurry"
import "curry2go/LogicExamples"
import "curry2go/Network"
import "curry2go/Perceptron"
import "curry2go/Prelude"

var func_names []string = []string{ "testNetwork", "testNetwork._#lambda1", "testLogic", "testNand11", "testNand11._#selFP2#example11", "testNand11._#selFP2#example11_CASE0", "testNand11._#selFP2#example11_CASE1", "testNand11._#selFP2#example11_CASE2", "testNand11._#selFP2#example11_CASE3", "testNetworkWs", "testNetworkWs._#lambda2" }

func TestPerceptron__CREATE_testNetwork( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestPerceptron_testNetwork, & func_names[ 0 ], 1, -1, args... )
    return( root )
}

func TestPerceptron__CREATE_testNetworkDot_Us_Hash_lambda1( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestPerceptron_testNetworkDot_Us_Hash_lambda1, & func_names[ 1 ], 2, -1, args... )
    return( root )
}

func TestPerceptron__CREATE_testLogic( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestPerceptron_testLogic, & func_names[ 2 ], 0, -1, args... )
    return( root )
}

func TestPerceptron__CREATE_testNand11( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestPerceptron_testNand11, & func_names[ 3 ], 0, -1, args... )
    return( root )
}

func TestPerceptron__CREATE_testNand11Dot_Us_Hash_selFP2Hash_example11( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestPerceptron_testNand11Dot_Us_Hash_selFP2Hash_example11, & func_names[ 4 ], 1, 0, args... )
    return( root )
}

func TestPerceptron__CREATE_testNand11Dot_Us_Hash_selFP2Hash_example11Us_CASE0( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestPerceptron_testNand11Dot_Us_Hash_selFP2Hash_example11Us_CASE0, & func_names[ 5 ], 1, 0, args... )
    return( root )
}

func TestPerceptron__CREATE_testNand11Dot_Us_Hash_selFP2Hash_example11Us_CASE1( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestPerceptron_testNand11Dot_Us_Hash_selFP2Hash_example11Us_CASE1, & func_names[ 6 ], 1, 0, args... )
    return( root )
}

func TestPerceptron__CREATE_testNand11Dot_Us_Hash_selFP2Hash_example11Us_CASE2( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestPerceptron_testNand11Dot_Us_Hash_selFP2Hash_example11Us_CASE2, & func_names[ 7 ], 1, 0, args... )
    return( root )
}

func TestPerceptron__CREATE_testNand11Dot_Us_Hash_selFP2Hash_example11Us_CASE3( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestPerceptron_testNand11Dot_Us_Hash_selFP2Hash_example11Us_CASE3, & func_names[ 8 ], 2, 0, args... )
    return( root )
}

func TestPerceptron__CREATE_testNetworkWs( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestPerceptron_testNetworkWs, & func_names[ 9 ], 2, -1, args... )
    return( root )
}

func TestPerceptron__CREATE_testNetworkWsDot_Us_Hash_lambda2( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestPerceptron_testNetworkWsDot_Us_Hash_lambda2, & func_names[ 10 ], 3, -1, args... )
    return( root )
}

func TestPerceptron_testNetwork( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    x1 = root.Children[ 0 ]
    Prelude.Prelude__CREATE_filter( root, TestPerceptron__CREATE_testNetworkDot_Us_Hash_lambda1( root.NewNode(  ), x1 ), Network.Network__CREATE_exs( root.NewNode(  ), x1 ) )
    return
}

func TestPerceptron_testNetworkDot_Us_Hash_lambda1( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    Prelude.Prelude__CREATE_apply( root, Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_Us_implHash_Slash_Eq_Hash_PreludeDot_EqHash_LSb_RSb_Hash_0Hash_Hash_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_instHash_PreludeDot_EqHash_PreludeDot_FloatHash_( root.NewNode(  ) ) ), Perceptron.Perceptron__CREATE_eval( root.NewNode(  ), x1, Network.Network__CREATE_inputs( root.NewNode(  ), x2 ) ) ), Network.Network__CREATE_expectation( root.NewNode(  ), x2 ) )
    return
}

func TestPerceptron_testLogic( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Prelude.Prelude__CREATE_Add_Add_( root, TestPerceptron__CREATE_testNetwork( root.NewNode(  ), LogicExamples.LogicExamples__CREATE_andNetwork( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_Add_Add_( root.NewNode(  ), TestPerceptron__CREATE_testNetwork( root.NewNode(  ), LogicExamples.LogicExamples__CREATE_orNetwork( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_Add_Add_( root.NewNode(  ), TestPerceptron__CREATE_testNetwork( root.NewNode(  ), LogicExamples.LogicExamples__CREATE_notNetwork( root.NewNode(  ) ) ), TestPerceptron__CREATE_testNetwork( root.NewNode(  ), LogicExamples.LogicExamples__CREATE_nandNetwork( root.NewNode(  ) ) ) ) ) )
    return
}

func TestPerceptron_testNand11( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    x1 = Network.Network__CREATE_exs( root.NewNode(  ), LogicExamples.LogicExamples__CREATE_nandNetwork( root.NewNode(  ) ) )
    x2 = TestPerceptron__CREATE_testNand11Dot_Us_Hash_selFP2Hash_example11( root.NewNode(  ), x1 )
    Perceptron.Perceptron__CREATE_networkDotProduct( root, LogicExamples.LogicExamples__CREATE_nandNetwork( root.NewNode(  ) ), Network.Network__CREATE_inputs( root.NewNode(  ), x2 ), Network.Network__CREATE_weights( root.NewNode(  ), LogicExamples.LogicExamples__CREATE_nandNetwork( root.NewNode(  ) ) ) )
    return
}

func TestPerceptron_testNand11Dot_Us_Hash_selFP2Hash_example11( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    x1 = root.Children[ 0 ]
    switch x1.GetConstructor(  ){
        case -1:
            if( task.IsBound( x1 ) ){
                task.ToHnf( x1 )
                return
            }else {
            }
            x1.SetTrLock( task.GetId(  ), gocurry.ChoiceCreate( task.NewNode(  ), Prelude.Prelude__CREATE_LSb_RSb_( task.NewNode(  ) ), Prelude.Prelude__CREATE_Col_( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) ) )
            return
        case 0:
            gocurry.ExemptCreate( root )
            return
        case 1:
            var x3 *gocurry.Node
            x3 = x1.Children[ 1 ]
            TestPerceptron__CREATE_testNand11Dot_Us_Hash_selFP2Hash_example11Us_CASE0( root, x3 )
            return
    }
}

func TestPerceptron_testNand11Dot_Us_Hash_selFP2Hash_example11Us_CASE0( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x3 *gocurry.Node
    x3 = root.Children[ 0 ]
    switch x3.GetConstructor(  ){
        case -1:
            if( task.IsBound( x3 ) ){
                task.ToHnf( x3 )
                return
            }else {
            }
            x3.SetTrLock( task.GetId(  ), gocurry.ChoiceCreate( task.NewNode(  ), Prelude.Prelude__CREATE_LSb_RSb_( task.NewNode(  ) ), Prelude.Prelude__CREATE_Col_( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) ) )
            return
        case 0:
            gocurry.ExemptCreate( root )
            return
        case 1:
            var x5 *gocurry.Node
            x5 = x3.Children[ 1 ]
            TestPerceptron__CREATE_testNand11Dot_Us_Hash_selFP2Hash_example11Us_CASE1( root, x5 )
            return
    }
}

func TestPerceptron_testNand11Dot_Us_Hash_selFP2Hash_example11Us_CASE1( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x5 *gocurry.Node
    x5 = root.Children[ 0 ]
    switch x5.GetConstructor(  ){
        case -1:
            if( task.IsBound( x5 ) ){
                task.ToHnf( x5 )
                return
            }else {
            }
            x5.SetTrLock( task.GetId(  ), gocurry.ChoiceCreate( task.NewNode(  ), Prelude.Prelude__CREATE_LSb_RSb_( task.NewNode(  ) ), Prelude.Prelude__CREATE_Col_( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) ) )
            return
        case 0:
            gocurry.ExemptCreate( root )
            return
        case 1:
            var x7 *gocurry.Node
            x7 = x5.Children[ 1 ]
            TestPerceptron__CREATE_testNand11Dot_Us_Hash_selFP2Hash_example11Us_CASE2( root, x7 )
            return
    }
}

func TestPerceptron_testNand11Dot_Us_Hash_selFP2Hash_example11Us_CASE2( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x7 *gocurry.Node
    x7 = root.Children[ 0 ]
    switch x7.GetConstructor(  ){
        case -1:
            if( task.IsBound( x7 ) ){
                task.ToHnf( x7 )
                return
            }else {
            }
            x7.SetTrLock( task.GetId(  ), gocurry.ChoiceCreate( task.NewNode(  ), Prelude.Prelude__CREATE_LSb_RSb_( task.NewNode(  ) ), Prelude.Prelude__CREATE_Col_( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) ) )
            return
        case 0:
            gocurry.ExemptCreate( root )
            return
        case 1:
            var x8 *gocurry.Node
            var x9 *gocurry.Node
            x8 = x7.Children[ 0 ]
            x9 = x7.Children[ 1 ]
            TestPerceptron__CREATE_testNand11Dot_Us_Hash_selFP2Hash_example11Us_CASE3( root, x9, x8 )
            return
    }
}

func TestPerceptron_testNand11Dot_Us_Hash_selFP2Hash_example11Us_CASE3( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x9 *gocurry.Node
    var x8 *gocurry.Node
    x9 = root.Children[ 0 ]
    x8 = root.Children[ 1 ]
    switch x9.GetConstructor(  ){
        case -1:
            if( task.IsBound( x9 ) ){
                task.ToHnf( x9 )
                return
            }else {
            }
            x9.SetTrLock( task.GetId(  ), gocurry.ChoiceCreate( task.NewNode(  ), Prelude.Prelude__CREATE_LSb_RSb_( task.NewNode(  ) ), Prelude.Prelude__CREATE_Col_( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) ) )
            return
        case 0:
            gocurry.RedirectCreate( root, x8 )
            return
        case 1:
            gocurry.ExemptCreate( root )
            return
    }
}

func TestPerceptron_testNetworkWs( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    Prelude.Prelude__CREATE_filter( root, TestPerceptron__CREATE_testNetworkWsDot_Us_Hash_lambda2( root.NewNode(  ), x1, x2 ), Network.Network__CREATE_exs( root.NewNode(  ), x1 ) )
    return
}

func TestPerceptron_testNetworkWsDot_Us_Hash_lambda2( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    var x3 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    x3 = root.Children[ 2 ]
    Prelude.Prelude__CREATE_apply( root, Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_Us_implHash_Slash_Eq_Hash_PreludeDot_EqHash_LSb_RSb_Hash_0Hash_Hash_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_instHash_PreludeDot_EqHash_PreludeDot_FloatHash_( root.NewNode(  ) ) ), Perceptron.Perceptron__CREATE_eval2( root.NewNode(  ), x1, x2, Network.Network__CREATE_inputs( root.NewNode(  ), x3 ) ) ), Network.Network__CREATE_expectation( root.NewNode(  ), x3 ) )
    return
}

