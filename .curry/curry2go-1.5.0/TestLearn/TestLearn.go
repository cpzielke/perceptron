package TestLearn

import "gocurry"
import "curry2go/Algorithms"
import "curry2go/Learn"
import "curry2go/LogicExamples"
import "curry2go/Network"
import "curry2go/Perceptron"
import "curry2go/Prelude"

var func_names []string = []string{ "testLearn", "testInitialWeights", "testDeltaOut", "testDeltaOut2", "testDeltaNeuron", "testLayer", "testDeltaInner", "testBackPropagate", "testBackPropagate._#selFP2#example11", "testBackPropagate._#selFP2#example11_CASE0", "testBackPropagate._#selFP2#example11_CASE1", "testBackPropagate._#selFP2#example11_CASE2", "testBackPropagate._#selFP2#example11_CASE3", "testGD", "testLE", "testLEs", "testLEs._#lambda1", "testEpic", "testGate", "testLEgate", "testLearnGate", "testNetwork", "networkWeights", "testNetExample", "testFFnet", "testBPnet", "testGDnet", "testLEnet", "testEpicNet", "testLearnNet" }

func TestLearn__CREATE_testLearn( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testLearn, & func_names[ 0 ], 0, -1, args... )
    return( root )
}

func TestLearn__CREATE_testInitialWeights( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testInitialWeights, & func_names[ 1 ], 0, -1, args... )
    return( root )
}

func TestLearn__CREATE_testDeltaOut( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testDeltaOut, & func_names[ 2 ], 0, -1, args... )
    return( root )
}

func TestLearn__CREATE_testDeltaOut2( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testDeltaOut2, & func_names[ 3 ], 1, -1, args... )
    return( root )
}

func TestLearn__CREATE_testDeltaNeuron( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testDeltaNeuron, & func_names[ 4 ], 0, -1, args... )
    return( root )
}

func TestLearn__CREATE_testLayer( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testLayer, & func_names[ 5 ], 0, -1, args... )
    return( root )
}

func TestLearn__CREATE_testDeltaInner( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testDeltaInner, & func_names[ 6 ], 0, -1, args... )
    return( root )
}

func TestLearn__CREATE_testBackPropagate( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testBackPropagate, & func_names[ 7 ], 0, -1, args... )
    return( root )
}

func TestLearn__CREATE_testBackPropagateDot_Us_Hash_selFP2Hash_example11( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testBackPropagateDot_Us_Hash_selFP2Hash_example11, & func_names[ 8 ], 1, 0, args... )
    return( root )
}

func TestLearn__CREATE_testBackPropagateDot_Us_Hash_selFP2Hash_example11Us_CASE0( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testBackPropagateDot_Us_Hash_selFP2Hash_example11Us_CASE0, & func_names[ 9 ], 1, 0, args... )
    return( root )
}

func TestLearn__CREATE_testBackPropagateDot_Us_Hash_selFP2Hash_example11Us_CASE1( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testBackPropagateDot_Us_Hash_selFP2Hash_example11Us_CASE1, & func_names[ 10 ], 1, 0, args... )
    return( root )
}

func TestLearn__CREATE_testBackPropagateDot_Us_Hash_selFP2Hash_example11Us_CASE2( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testBackPropagateDot_Us_Hash_selFP2Hash_example11Us_CASE2, & func_names[ 11 ], 1, 0, args... )
    return( root )
}

func TestLearn__CREATE_testBackPropagateDot_Us_Hash_selFP2Hash_example11Us_CASE3( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testBackPropagateDot_Us_Hash_selFP2Hash_example11Us_CASE3, & func_names[ 12 ], 2, 0, args... )
    return( root )
}

func TestLearn__CREATE_testGD( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testGD, & func_names[ 13 ], 0, -1, args... )
    return( root )
}

func TestLearn__CREATE_testLE( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testLE, & func_names[ 14 ], 0, -1, args... )
    return( root )
}

func TestLearn__CREATE_testLEs( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testLEs, & func_names[ 15 ], 1, -1, args... )
    return( root )
}

func TestLearn__CREATE_testLEsDot_Us_Hash_lambda1( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testLEsDot_Us_Hash_lambda1, & func_names[ 16 ], 2, -1, args... )
    return( root )
}

func TestLearn__CREATE_testEpic( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testEpic, & func_names[ 17 ], 2, -1, args... )
    return( root )
}

func TestLearn__CREATE_testGate( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testGate, & func_names[ 18 ], 3, -1, args... )
    return( root )
}

func TestLearn__CREATE_testLEgate( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testLEgate, & func_names[ 19 ], 2, -1, args... )
    return( root )
}

func TestLearn__CREATE_testLearnGate( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testLearnGate, & func_names[ 20 ], 2, -1, args... )
    return( root )
}

func TestLearn__CREATE_testNetwork( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testNetwork, & func_names[ 21 ], 0, -1, args... )
    return( root )
}

func TestLearn__CREATE_networkWeights( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_networkWeights, & func_names[ 22 ], 0, -1, args... )
    return( root )
}

func TestLearn__CREATE_testNetExample( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testNetExample, & func_names[ 23 ], 0, -1, args... )
    return( root )
}

func TestLearn__CREATE_testFFnet( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testFFnet, & func_names[ 24 ], 0, -1, args... )
    return( root )
}

func TestLearn__CREATE_testBPnet( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testBPnet, & func_names[ 25 ], 0, -1, args... )
    return( root )
}

func TestLearn__CREATE_testGDnet( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testGDnet, & func_names[ 26 ], 0, -1, args... )
    return( root )
}

func TestLearn__CREATE_testLEnet( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testLEnet, & func_names[ 27 ], 0, -1, args... )
    return( root )
}

func TestLearn__CREATE_testEpicNet( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testEpicNet, & func_names[ 28 ], 0, -1, args... )
    return( root )
}

func TestLearn__CREATE_testLearnNet( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, TestLearn_testLearnNet, & func_names[ 29 ], 0, -1, args... )
    return( root )
}

func TestLearn_testLearn( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Learn.Learn__CREATE_learn( root, gocurry.IntLitCreate( root.NewNode(  ), 1 ), LogicExamples.LogicExamples__CREATE_andNetwork( root.NewNode(  ) ), Network.Network__CREATE_weights( root.NewNode(  ), LogicExamples.LogicExamples__CREATE_andNetwork( root.NewNode(  ) ) ) )
    return
}

func TestLearn_testInitialWeights( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Learn.Learn__CREATE_initialWeights( root, LogicExamples.LogicExamples__CREATE_nandNetwork( root.NewNode(  ) ), gocurry.IntLitCreate( root.NewNode(  ), 2 ) )
    return
}

func TestLearn_testDeltaOut( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Learn.Learn__CREATE_deltaOut( root, Prelude.Prelude__CREATE_head( root.NewNode(  ), Network.Network__CREATE_layers( root.NewNode(  ), LogicExamples.LogicExamples__CREATE_andNetwork( root.NewNode(  ) ) ) ), Network.Network__CREATE_expectation( root.NewNode(  ), Prelude.Prelude__CREATE_head( root.NewNode(  ), Network.Network__CREATE_exs( root.NewNode(  ), LogicExamples.LogicExamples__CREATE_andNetwork( root.NewNode(  ) ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) )
    return
}

func TestLearn_testDeltaOut2( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    x1 = root.Children[ 0 ]
    Learn.Learn__CREATE_deltaOut( root, Prelude.Prelude__CREATE_head( root.NewNode(  ), Network.Network__CREATE_layers( root.NewNode(  ), LogicExamples.LogicExamples__CREATE_andNetwork( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), x1 )
    return
}

func TestLearn_testDeltaNeuron( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Prelude.Prelude__CREATE_Us_implHash_Eq_Eq_Hash_PreludeDot_EqHash_PreludeDot_FloatHash_( root, Learn.Learn__CREATE_deltaNeuron( root.NewNode(  ), gocurry.IntLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 2 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 3 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 4 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 5 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 6 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 7 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 8 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 2 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 3 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) ), gocurry.FloatLitCreate( root.NewNode(  ), 17 ) )
    return
}

func TestLearn_testLayer( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Network.Network__CREATE_Layer( root, gocurry.IntLitCreate( root.NewNode(  ), 3 ), Algorithms.Algorithms__CREATE_classifier( root.NewNode(  ) ) )
    return
}

func TestLearn_testDeltaInner( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Prelude.Prelude__CREATE_Us_implHash_Eq_Eq_Hash_PreludeDot_EqHash_LSb_RSb_Hash_0Hash_Hash_( root, Prelude.Prelude__CREATE_Us_instHash_PreludeDot_EqHash_PreludeDot_FloatHash_( root.NewNode(  ) ), Learn.Learn__CREATE_deltaInner( root.NewNode(  ), gocurry.IntLitCreate( root.NewNode(  ), 0 ), TestLearn__CREATE_testLayer( root.NewNode(  ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 2 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 3 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 4 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 5 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 6 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 7 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 8 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 2 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 3 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 10 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 20 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 30 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 17 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 22 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 27 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) ) )
    return
}

func TestLearn_testBackPropagate( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    x1 = Network.Network__CREATE_exs( root.NewNode(  ), LogicExamples.LogicExamples__CREATE_nandNetwork( root.NewNode(  ) ) )
    x2 = TestLearn__CREATE_testBackPropagateDot_Us_Hash_selFP2Hash_example11( root.NewNode(  ), x1 )
    Learn.Learn__CREATE_backPropagate( root, Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_reverse( root.NewNode(  ) ), Network.Network__CREATE_layers( root.NewNode(  ), LogicExamples.LogicExamples__CREATE_nandNetwork( root.NewNode(  ) ) ) ), Network.Network__CREATE_expectation( root.NewNode(  ), x2 ), Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_reverse( root.NewNode(  ) ), Network.Network__CREATE_weights( root.NewNode(  ), LogicExamples.LogicExamples__CREATE_nandNetwork( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) ) )
    return
}

func TestLearn_testBackPropagateDot_Us_Hash_selFP2Hash_example11( task *gocurry.Task )(  ){
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
            TestLearn__CREATE_testBackPropagateDot_Us_Hash_selFP2Hash_example11Us_CASE0( root, x3 )
            return
    }
}

func TestLearn_testBackPropagateDot_Us_Hash_selFP2Hash_example11Us_CASE0( task *gocurry.Task )(  ){
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
            TestLearn__CREATE_testBackPropagateDot_Us_Hash_selFP2Hash_example11Us_CASE1( root, x5 )
            return
    }
}

func TestLearn_testBackPropagateDot_Us_Hash_selFP2Hash_example11Us_CASE1( task *gocurry.Task )(  ){
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
            TestLearn__CREATE_testBackPropagateDot_Us_Hash_selFP2Hash_example11Us_CASE2( root, x7 )
            return
    }
}

func TestLearn_testBackPropagateDot_Us_Hash_selFP2Hash_example11Us_CASE2( task *gocurry.Task )(  ){
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
            TestLearn__CREATE_testBackPropagateDot_Us_Hash_selFP2Hash_example11Us_CASE3( root, x9, x8 )
            return
    }
}

func TestLearn_testBackPropagateDot_Us_Hash_selFP2Hash_example11Us_CASE3( task *gocurry.Task )(  ){
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

func TestLearn_testGD( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Learn.Learn__CREATE_gradientDescent( root, Network.Network__CREATE_learningRate( root.NewNode(  ), LogicExamples.LogicExamples__CREATE_nandNetwork( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_reverse( root.NewNode(  ) ), Network.Network__CREATE_weights( root.NewNode(  ), LogicExamples.LogicExamples__CREATE_nandNetwork( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) )
    return
}

func TestLearn_testLE( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Learn.Learn__CREATE_learnExample( root, LogicExamples.LogicExamples__CREATE_nandNetwork( root.NewNode(  ) ), Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_reverse( root.NewNode(  ) ), Network.Network__CREATE_weights( root.NewNode(  ), LogicExamples.LogicExamples__CREATE_nandNetwork( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_head( root.NewNode(  ), Network.Network__CREATE_exs( root.NewNode(  ), LogicExamples.LogicExamples__CREATE_nandNetwork( root.NewNode(  ) ) ) ) )
    return
}

func TestLearn_testLEs( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    x1 = root.Children[ 0 ]
    Prelude.Prelude__CREATE_filter( root, TestLearn__CREATE_testLEsDot_Us_Hash_lambda1( root.NewNode(  ), x1 ), Network.Network__CREATE_exs( root.NewNode(  ), x1 ) )
    return
}

func TestLearn_testLEsDot_Us_Hash_lambda1( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    Prelude.Prelude__CREATE_apply( root, Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_Us_implHash_Slash_Eq_Hash_PreludeDot_EqHash_LSb_RSb_Hash_0Hash_Hash_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_instHash_PreludeDot_EqHash_LSb_RSb_Hash_0Hash_Hash_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_instHash_PreludeDot_EqHash_LSb_RSb_Hash_0Hash_Hash_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_instHash_PreludeDot_EqHash_PreludeDot_FloatHash_( root.NewNode(  ) ) ) ) ), Learn.Learn__CREATE_learnExample( root.NewNode(  ), x1, Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_reverse( root.NewNode(  ) ), Network.Network__CREATE_weights( root.NewNode(  ), x1 ) ), x2 ) ), Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_reverse( root.NewNode(  ) ), Network.Network__CREATE_weights( root.NewNode(  ), x1 ) ) )
    return
}

func TestLearn_testEpic( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    Learn.Learn__CREATE_learnEpic( root, x1, x2 )
    return
}

func TestLearn_testGate( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    var x3 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    x3 = root.Children[ 2 ]
    Perceptron.Perceptron__CREATE_eval2( root, x1, x3, x2 )
    return
}

func TestLearn_testLEgate( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    Learn.Learn__CREATE_learnExample( root, x1, x2, Prelude.Prelude__CREATE_head( root.NewNode(  ), Network.Network__CREATE_exs( root.NewNode(  ), x1 ) ) )
    return
}

func TestLearn_testLearnGate( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    Learn.Learn__CREATE_learn( root, gocurry.IntLitCreate( root.NewNode(  ), 1 ), x1, x2 )
    return
}

func TestLearn_testNetwork( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    LogicExamples.LogicExamples__CREATE_xorNetwork( root )
    return
}

func TestLearn_networkWeights( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Prelude.Prelude__CREATE_Col_( root, Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0.2 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0.2 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0.2 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0.3 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0.3 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0.3 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0.4 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0.4 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0.4 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) )
    return
}

func TestLearn_testNetExample( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Prelude.Prelude__CREATE_head( root, Network.Network__CREATE_exs( root.NewNode(  ), TestLearn__CREATE_testNetwork( root.NewNode(  ) ) ) )
    return
}

func TestLearn_testFFnet( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Perceptron.Perceptron__CREATE_feedForward( root, TestLearn__CREATE_testNetwork( root.NewNode(  ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), TestLearn__CREATE_networkWeights( root.NewNode(  ) ) )
    return
}

func TestLearn_testBPnet( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Learn.Learn__CREATE_backPropagate( root, Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_reverse( root.NewNode(  ) ), Network.Network__CREATE_layers( root.NewNode(  ), TestLearn__CREATE_testNetwork( root.NewNode(  ) ) ) ), Network.Network__CREATE_expectation( root.NewNode(  ), TestLearn__CREATE_testNetExample( root.NewNode(  ) ) ), TestLearn__CREATE_networkWeights( root.NewNode(  ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0.574442516811659 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0.5986876601124521 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) ) )
    return
}

func TestLearn_testGDnet( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Learn.Learn__CREATE_gradientDescent( root, Network.Network__CREATE_learningRate( root.NewNode(  ), TestLearn__CREATE_testNetwork( root.NewNode(  ) ) ), TestLearn__CREATE_networkWeights( root.NewNode(  ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0.574442516811659 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0.5986876601124521 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), Prelude.Prelude__CREATE_Col_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) ) )
    return
}

func TestLearn_testLEnet( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Learn.Learn__CREATE_learnExample( root, TestLearn__CREATE_testNetwork( root.NewNode(  ) ), TestLearn__CREATE_networkWeights( root.NewNode(  ) ), TestLearn__CREATE_testNetExample( root.NewNode(  ) ) )
    return
}

func TestLearn_testEpicNet( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Learn.Learn__CREATE_learnEpic( root, TestLearn__CREATE_testNetwork( root.NewNode(  ) ), TestLearn__CREATE_networkWeights( root.NewNode(  ) ) )
    return
}

func TestLearn_testLearnNet( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Learn.Learn__CREATE_learn( root, gocurry.IntLitCreate( root.NewNode(  ), 1000 ), TestLearn__CREATE_testNetwork( root.NewNode(  ) ), TestLearn__CREATE_networkWeights( root.NewNode(  ) ) )
    return
}

