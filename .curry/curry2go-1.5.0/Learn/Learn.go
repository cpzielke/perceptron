package Learn

import "gocurry"
import "curry2go/Algorithms"
import "curry2go/Network"
import "curry2go/Perceptron"
import "curry2go/Prelude"

var func_names []string = []string{ "deltaOut", "deltaOut_CASE1", "deltaOut_LET2", "deltaOut_CASE0", "deltaNeuron", "deltaInner", "deltaInner_LET0", "backPropagate", "backPropagate_CASE0", "backPropagate.backProp.24", "backPropagate.backProp.24_CASE0", "backPropagate.backProp.24_CASE1", "backPropagate.backProp.24_LET2", "adjustNeuron", "adjustNeuron_CASE0", "adjustNeuron_CASE1", "adjustLayer", "adjustLayer_CASE0", "adjustLayer_LET1", "adjustLayer_COMPLEXCASE2", "gradientDescent", "gradientDescent_CASE0", "gradientDescent_CASE1", "gradientDescent_CASE2", "gradientDescent_LET3", "gradientDescent_COMPLEXCASE4", "learnExample", "learnExampleNetwork", "learnEpic", "learn", "learn_COMPLEXCASE0", "initialWeights", "initialWeights.initLWs.64", "initialWeights.initLWs.64_LET0", "initialWeights.initLWs.64.initNWs.69", "initialWeights.initLWs.64.initNWs.69_COMPLEXCASE0", "initialWeights.initLWs.64.initNWs.69.initNWs'.72", "initialWeights.initLWs.64.initNWs.69.initNWs'.72.initWs.74" }

func Learn__CREATE_deltaOut( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_deltaOut, & func_names[ 0 ], 3, 1, args... )
    return( root )
}

func Learn__CREATE_deltaOutUs_CASE1( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_deltaOutUs_CASE1, & func_names[ 1 ], 4, 0, args... )
    return( root )
}

func Learn__CREATE_deltaOutUs_LET2( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_deltaOutUs_LET2, & func_names[ 2 ], 5, -1, args... )
    return( root )
}

func Learn__CREATE_deltaOutUs_CASE0( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_deltaOutUs_CASE0, & func_names[ 3 ], 1, 0, args... )
    return( root )
}

func Learn__CREATE_deltaNeuron( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_deltaNeuron, & func_names[ 4 ], 3, -1, args... )
    return( root )
}

func Learn__CREATE_deltaInner( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_deltaInner, & func_names[ 5 ], 5, 4, args... )
    return( root )
}

func Learn__CREATE_deltaInnerUs_LET0( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_deltaInnerUs_LET0, & func_names[ 6 ], 6, -1, args... )
    return( root )
}

func Learn__CREATE_backPropagate( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_backPropagate, & func_names[ 7 ], 4, 0, args... )
    return( root )
}

func Learn__CREATE_backPropagateUs_CASE0( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_backPropagateUs_CASE0, & func_names[ 8 ], 5, 0, args... )
    return( root )
}

func Learn__CREATE_backPropagateDot_backPropDot_24( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_backPropagateDot_backPropDot_24, & func_names[ 9 ], 4, 0, args... )
    return( root )
}

func Learn__CREATE_backPropagateDot_backPropDot_24Us_CASE0( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_backPropagateDot_backPropDot_24Us_CASE0, & func_names[ 10 ], 5, 0, args... )
    return( root )
}

func Learn__CREATE_backPropagateDot_backPropDot_24Us_CASE1( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_backPropagateDot_backPropDot_24Us_CASE1, & func_names[ 11 ], 6, 0, args... )
    return( root )
}

func Learn__CREATE_backPropagateDot_backPropDot_24Us_LET2( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_backPropagateDot_backPropDot_24Us_LET2, & func_names[ 12 ], 7, -1, args... )
    return( root )
}

func Learn__CREATE_adjustNeuron( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_adjustNeuron, & func_names[ 13 ], 4, 1, args... )
    return( root )
}

func Learn__CREATE_adjustNeuronUs_CASE0( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_adjustNeuronUs_CASE0, & func_names[ 14 ], 5, 0, args... )
    return( root )
}

func Learn__CREATE_adjustNeuronUs_CASE1( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_adjustNeuronUs_CASE1, & func_names[ 15 ], 4, 0, args... )
    return( root )
}

func Learn__CREATE_adjustLayer( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_adjustLayer, & func_names[ 16 ], 4, 1, args... )
    return( root )
}

func Learn__CREATE_adjustLayerUs_CASE0( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_adjustLayerUs_CASE0, & func_names[ 17 ], 5, 0, args... )
    return( root )
}

func Learn__CREATE_adjustLayerUs_LET1( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_adjustLayerUs_LET1, & func_names[ 18 ], 6, -1, args... )
    return( root )
}

func Learn__CREATE_adjustLayerUs_COMPLEXCASE2( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_adjustLayerUs_COMPLEXCASE2, & func_names[ 19 ], 6, 5, args... )
    return( root )
}

func Learn__CREATE_gradientDescent( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_gradientDescent, & func_names[ 20 ], 4, 1, args... )
    return( root )
}

func Learn__CREATE_gradientDescentUs_CASE0( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_gradientDescentUs_CASE0, & func_names[ 21 ], 5, 0, args... )
    return( root )
}

func Learn__CREATE_gradientDescentUs_CASE1( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_gradientDescentUs_CASE1, & func_names[ 22 ], 5, 0, args... )
    return( root )
}

func Learn__CREATE_gradientDescentUs_CASE2( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_gradientDescentUs_CASE2, & func_names[ 23 ], 6, 0, args... )
    return( root )
}

func Learn__CREATE_gradientDescentUs_LET3( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_gradientDescentUs_LET3, & func_names[ 24 ], 7, -1, args... )
    return( root )
}

func Learn__CREATE_gradientDescentUs_COMPLEXCASE4( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_gradientDescentUs_COMPLEXCASE4, & func_names[ 25 ], 7, 6, args... )
    return( root )
}

func Learn__CREATE_learnExample( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_learnExample, & func_names[ 26 ], 3, -1, args... )
    return( root )
}

func Learn__CREATE_learnExampleNetwork( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_learnExampleNetwork, & func_names[ 27 ], 2, -1, args... )
    return( root )
}

func Learn__CREATE_learnEpic( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_learnEpic, & func_names[ 28 ], 2, -1, args... )
    return( root )
}

func Learn__CREATE_learn( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_learn, & func_names[ 29 ], 3, -1, args... )
    return( root )
}

func Learn__CREATE_learnUs_COMPLEXCASE0( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_learnUs_COMPLEXCASE0, & func_names[ 30 ], 5, 4, args... )
    return( root )
}

func Learn__CREATE_initialWeights( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_initialWeights, & func_names[ 31 ], 2, -1, args... )
    return( root )
}

func Learn__CREATE_initialWeightsDot_initLWsDot_64( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_initialWeightsDot_initLWsDot_64, & func_names[ 32 ], 2, 1, args... )
    return( root )
}

func Learn__CREATE_initialWeightsDot_initLWsDot_64Us_LET0( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_initialWeightsDot_initLWsDot_64Us_LET0, & func_names[ 33 ], 3, -1, args... )
    return( root )
}

func Learn__CREATE_initialWeightsDot_initLWsDot_64Dot_initNWsDot_69( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_initialWeightsDot_initLWsDot_64Dot_initNWsDot_69, & func_names[ 34 ], 6, -1, args... )
    return( root )
}

func Learn__CREATE_initialWeightsDot_initLWsDot_64Dot_initNWsDot_69Us_COMPLEXCASE0( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_initialWeightsDot_initLWsDot_64Dot_initNWsDot_69Us_COMPLEXCASE0, & func_names[ 35 ], 7, 6, args... )
    return( root )
}

func Learn__CREATE_initialWeightsDot_initLWsDot_64Dot_initNWsDot_69Dot_initNWsSQuote_Dot_72( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_initialWeightsDot_initLWsDot_64Dot_initNWsDot_69Dot_initNWsSQuote_Dot_72, & func_names[ 36 ], 6, -1, args... )
    return( root )
}

func Learn__CREATE_initialWeightsDot_initLWsDot_64Dot_initNWsDot_69Dot_initNWsSQuote_Dot_72Dot_initWsDot_74( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Learn_initialWeightsDot_initLWsDot_64Dot_initNWsDot_69Dot_initNWsSQuote_Dot_72Dot_initWsDot_74, & func_names[ 37 ], 3, -1, args... )
    return( root )
}

func Learn_deltaOut( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    var x3 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    x3 = root.Children[ 2 ]
    switch x2.GetConstructor(  ){
        case -1:
            if( task.IsBound( x2 ) ){
                task.ToHnf( x2 )
                return
            }else {
            }
            x2.SetTrLock( task.GetId(  ), gocurry.ChoiceCreate( task.NewNode(  ), Prelude.Prelude__CREATE_LSb_RSb_( task.NewNode(  ) ), Prelude.Prelude__CREATE_Col_( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) ) )
            return
        case 0:
            Learn__CREATE_deltaOutUs_CASE0( root, x3 )
            return
        case 1:
            var x6 *gocurry.Node
            var x7 *gocurry.Node
            x6 = x2.Children[ 0 ]
            x7 = x2.Children[ 1 ]
            Learn__CREATE_deltaOutUs_CASE1( root, x3, x6, x7, x1 )
            return
    }
}

func Learn_deltaOutUs_CASE1( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x3 *gocurry.Node
    var x6 *gocurry.Node
    var x7 *gocurry.Node
    var x1 *gocurry.Node
    x3 = root.Children[ 0 ]
    x6 = root.Children[ 1 ]
    x7 = root.Children[ 2 ]
    x1 = root.Children[ 3 ]
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
            var x8 *gocurry.Node
            var x9 *gocurry.Node
            x8 = x3.Children[ 0 ]
            x9 = x3.Children[ 1 ]
            Learn__CREATE_deltaOutUs_LET2( root, x6, x8, x7, x9, x1 )
            return
    }
}

func Learn_deltaOutUs_LET2( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x6 *gocurry.Node
    var x8 *gocurry.Node
    var x7 *gocurry.Node
    var x9 *gocurry.Node
    var x1 *gocurry.Node
    var x10 *gocurry.Node
    x6 = root.Children[ 0 ]
    x8 = root.Children[ 1 ]
    x7 = root.Children[ 2 ]
    x9 = root.Children[ 3 ]
    x1 = root.Children[ 4 ]
    x10 = Algorithms.Algorithms__CREATE_activationUs_der( root.NewNode(  ), Network.Network__CREATE_algorithm( root.NewNode(  ), x1 ) )
    Prelude.Prelude__CREATE_Col_( root, Prelude.Prelude__CREATE_Us_implHash_Mul_Hash_PreludeDot_NumHash_PreludeDot_FloatHash_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_implHash_Sub_Hash_PreludeDot_NumHash_PreludeDot_FloatHash_( root.NewNode(  ), x8, x6 ), Prelude.Prelude__CREATE_apply( root.NewNode(  ), x10, x8 ) ), Learn__CREATE_deltaOut( root.NewNode(  ), x1, x7, x9 ) )
    return
}

func Learn_deltaOutUs_CASE0( task *gocurry.Task )(  ){
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
            Prelude.Prelude__CREATE_LSb_RSb_( root )
            return
        case 1:
            gocurry.ExemptCreate( root )
            return
    }
}

func Learn_deltaNeuron( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    var x3 *gocurry.Node
    var x4 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    x3 = root.Children[ 2 ]
    x4 = Prelude.Prelude__CREATE_map( root.NewNode(  ), Prelude.Prelude__CREATE_flip( root.NewNode(  ), Prelude.Prelude__CREATE_Excl_Excl_( root.NewNode(  ) ), x1 ), x2 )
    Perceptron.Perceptron__CREATE_dotProduct( root, x4, x3 )
    return
}

func Learn_deltaInner( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    var x3 *gocurry.Node
    var x4 *gocurry.Node
    var x5 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    x3 = root.Children[ 2 ]
    x4 = root.Children[ 3 ]
    x5 = root.Children[ 4 ]
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
            Prelude.Prelude__CREATE_LSb_RSb_( root )
            return
        case 1:
            var x6 *gocurry.Node
            var x7 *gocurry.Node
            x6 = x5.Children[ 0 ]
            x7 = x5.Children[ 1 ]
            Learn__CREATE_deltaInnerUs_LET0( root, x7, x2, x1, x3, x4, x6 )
            return
    }
}

func Learn_deltaInnerUs_LET0( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x7 *gocurry.Node
    var x2 *gocurry.Node
    var x1 *gocurry.Node
    var x3 *gocurry.Node
    var x4 *gocurry.Node
    var x6 *gocurry.Node
    var x8 *gocurry.Node
    var x9 *gocurry.Node
    x7 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    x1 = root.Children[ 2 ]
    x3 = root.Children[ 3 ]
    x4 = root.Children[ 4 ]
    x6 = root.Children[ 5 ]
    x8 = Algorithms.Algorithms__CREATE_activationUs_der( root.NewNode(  ), Network.Network__CREATE_algorithm( root.NewNode(  ), x2 ) )
    x9 = Prelude.Prelude__CREATE_Us_implHash_Mul_Hash_PreludeDot_NumHash_PreludeDot_FloatHash_( root.NewNode(  ), Learn__CREATE_deltaNeuron( root.NewNode(  ), x1, x3, x4 ), Prelude.Prelude__CREATE_apply( root.NewNode(  ), x8, x6 ) )
    Prelude.Prelude__CREATE_Col_( root, x9, Learn__CREATE_deltaInner( root.NewNode(  ), Prelude.Prelude__CREATE_Us_implHash_Add_Hash_PreludeDot_NumHash_PreludeDot_IntHash_( root.NewNode(  ), x1, gocurry.IntLitCreate( root.NewNode(  ), 1 ) ), x2, x3, x4, x7 ) )
    return
}

func Learn_backPropagate( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    var x3 *gocurry.Node
    var x4 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    x3 = root.Children[ 2 ]
    x4 = root.Children[ 3 ]
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
            var x5 *gocurry.Node
            var x6 *gocurry.Node
            x5 = x1.Children[ 0 ]
            x6 = x1.Children[ 1 ]
            Learn__CREATE_backPropagateUs_CASE0( root, x4, x6, x3, x5, x2 )
            return
    }
}

func Learn_backPropagateUs_CASE0( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x4 *gocurry.Node
    var x6 *gocurry.Node
    var x3 *gocurry.Node
    var x5 *gocurry.Node
    var x2 *gocurry.Node
    x4 = root.Children[ 0 ]
    x6 = root.Children[ 1 ]
    x3 = root.Children[ 2 ]
    x5 = root.Children[ 3 ]
    x2 = root.Children[ 4 ]
    switch x4.GetConstructor(  ){
        case -1:
            if( task.IsBound( x4 ) ){
                task.ToHnf( x4 )
                return
            }else {
            }
            x4.SetTrLock( task.GetId(  ), gocurry.ChoiceCreate( task.NewNode(  ), Prelude.Prelude__CREATE_LSb_RSb_( task.NewNode(  ) ), Prelude.Prelude__CREATE_Col_( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) ) )
            return
        case 0:
            gocurry.ExemptCreate( root )
            return
        case 1:
            var x7 *gocurry.Node
            var x8 *gocurry.Node
            x7 = x4.Children[ 0 ]
            x8 = x4.Children[ 1 ]
            Learn__CREATE_backPropagateDot_backPropDot_24( root, x6, x3, Learn__CREATE_deltaOut( root.NewNode(  ), x5, x2, x7 ), x8 )
            return
    }
}

func Learn_backPropagateDot_backPropDot_24( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    var x3 *gocurry.Node
    var x4 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    x3 = root.Children[ 2 ]
    x4 = root.Children[ 3 ]
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
            Prelude.Prelude__CREATE_Col_( root, x3, Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) )
            return
        case 1:
            var x5 *gocurry.Node
            var x6 *gocurry.Node
            x5 = x1.Children[ 0 ]
            x6 = x1.Children[ 1 ]
            Learn__CREATE_backPropagateDot_backPropDot_24Us_CASE0( root, x2, x4, x6, x5, x3 )
            return
    }
}

func Learn_backPropagateDot_backPropDot_24Us_CASE0( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x2 *gocurry.Node
    var x4 *gocurry.Node
    var x6 *gocurry.Node
    var x5 *gocurry.Node
    var x3 *gocurry.Node
    x2 = root.Children[ 0 ]
    x4 = root.Children[ 1 ]
    x6 = root.Children[ 2 ]
    x5 = root.Children[ 3 ]
    x3 = root.Children[ 4 ]
    switch x2.GetConstructor(  ){
        case -1:
            if( task.IsBound( x2 ) ){
                task.ToHnf( x2 )
                return
            }else {
            }
            x2.SetTrLock( task.GetId(  ), gocurry.ChoiceCreate( task.NewNode(  ), Prelude.Prelude__CREATE_LSb_RSb_( task.NewNode(  ) ), Prelude.Prelude__CREATE_Col_( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) ) )
            return
        case 0:
            gocurry.ExemptCreate( root )
            return
        case 1:
            var x7 *gocurry.Node
            var x8 *gocurry.Node
            x7 = x2.Children[ 0 ]
            x8 = x2.Children[ 1 ]
            Learn__CREATE_backPropagateDot_backPropDot_24Us_CASE1( root, x4, x6, x8, x5, x7, x3 )
            return
    }
}

func Learn_backPropagateDot_backPropDot_24Us_CASE1( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x4 *gocurry.Node
    var x6 *gocurry.Node
    var x8 *gocurry.Node
    var x5 *gocurry.Node
    var x7 *gocurry.Node
    var x3 *gocurry.Node
    x4 = root.Children[ 0 ]
    x6 = root.Children[ 1 ]
    x8 = root.Children[ 2 ]
    x5 = root.Children[ 3 ]
    x7 = root.Children[ 4 ]
    x3 = root.Children[ 5 ]
    switch x4.GetConstructor(  ){
        case -1:
            if( task.IsBound( x4 ) ){
                task.ToHnf( x4 )
                return
            }else {
            }
            x4.SetTrLock( task.GetId(  ), gocurry.ChoiceCreate( task.NewNode(  ), Prelude.Prelude__CREATE_LSb_RSb_( task.NewNode(  ) ), Prelude.Prelude__CREATE_Col_( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) ) )
            return
        case 0:
            gocurry.ExemptCreate( root )
            return
        case 1:
            var x9 *gocurry.Node
            var x10 *gocurry.Node
            x9 = x4.Children[ 0 ]
            x10 = x4.Children[ 1 ]
            Learn__CREATE_backPropagateDot_backPropDot_24Us_LET2( root, x6, x8, x10, x5, x7, x3, x9 )
            return
    }
}

func Learn_backPropagateDot_backPropDot_24Us_LET2( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x6 *gocurry.Node
    var x8 *gocurry.Node
    var x10 *gocurry.Node
    var x5 *gocurry.Node
    var x7 *gocurry.Node
    var x3 *gocurry.Node
    var x9 *gocurry.Node
    var x11 *gocurry.Node
    x6 = root.Children[ 0 ]
    x8 = root.Children[ 1 ]
    x10 = root.Children[ 2 ]
    x5 = root.Children[ 3 ]
    x7 = root.Children[ 4 ]
    x3 = root.Children[ 5 ]
    x9 = root.Children[ 6 ]
    x11 = Learn__CREATE_deltaInner( root.NewNode(  ), gocurry.IntLitCreate( root.NewNode(  ), 0 ), x5, x7, x3, x9 )
    Prelude.Prelude__CREATE_Col_( root, x3, Learn__CREATE_backPropagateDot_backPropDot_24( root.NewNode(  ), x6, x8, x11, x10 ) )
    return
}

func Learn_adjustNeuron( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    var x3 *gocurry.Node
    var x4 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    x3 = root.Children[ 2 ]
    x4 = root.Children[ 3 ]
    switch x2.GetConstructor(  ){
        case -1:
            if( task.IsBound( x2 ) ){
                task.ToHnf( x2 )
                return
            }else {
            }
            x2.SetTrLock( task.GetId(  ), gocurry.ChoiceCreate( task.NewNode(  ), Prelude.Prelude__CREATE_LSb_RSb_( task.NewNode(  ) ), Prelude.Prelude__CREATE_Col_( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) ) )
            return
        case 0:
            gocurry.ExemptCreate( root )
            return
        case 1:
            var x5 *gocurry.Node
            var x6 *gocurry.Node
            x5 = x2.Children[ 0 ]
            x6 = x2.Children[ 1 ]
            Learn__CREATE_adjustNeuronUs_CASE0( root, x3, x5, x1, x6, x4 )
            return
    }
}

func Learn_adjustNeuronUs_CASE0( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x3 *gocurry.Node
    var x5 *gocurry.Node
    var x1 *gocurry.Node
    var x6 *gocurry.Node
    var x4 *gocurry.Node
    x3 = root.Children[ 0 ]
    x5 = root.Children[ 1 ]
    x1 = root.Children[ 2 ]
    x6 = root.Children[ 3 ]
    x4 = root.Children[ 4 ]
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
            Learn__CREATE_adjustNeuronUs_CASE1( root, x6, x5, x1, x4 )
            return
        case 1:
            var x9 *gocurry.Node
            var x10 *gocurry.Node
            x9 = x3.Children[ 0 ]
            x10 = x3.Children[ 1 ]
            Prelude.Prelude__CREATE_Col_( root, Prelude.Prelude__CREATE_Us_implHash_Sub_Hash_PreludeDot_NumHash_PreludeDot_FloatHash_( root.NewNode(  ), x5, Prelude.Prelude__CREATE_Us_implHash_Mul_Hash_PreludeDot_NumHash_PreludeDot_FloatHash_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_implHash_Mul_Hash_PreludeDot_NumHash_PreludeDot_FloatHash_( root.NewNode(  ), x1, x4 ), x9 ) ), Learn__CREATE_adjustNeuron( root.NewNode(  ), x1, x6, x10, x4 ) )
            return
    }
}

func Learn_adjustNeuronUs_CASE1( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x6 *gocurry.Node
    var x5 *gocurry.Node
    var x1 *gocurry.Node
    var x4 *gocurry.Node
    x6 = root.Children[ 0 ]
    x5 = root.Children[ 1 ]
    x1 = root.Children[ 2 ]
    x4 = root.Children[ 3 ]
    switch x6.GetConstructor(  ){
        case -1:
            if( task.IsBound( x6 ) ){
                task.ToHnf( x6 )
                return
            }else {
            }
            x6.SetTrLock( task.GetId(  ), gocurry.ChoiceCreate( task.NewNode(  ), Prelude.Prelude__CREATE_LSb_RSb_( task.NewNode(  ) ), Prelude.Prelude__CREATE_Col_( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) ) )
            return
        case 0:
            Prelude.Prelude__CREATE_Col_( root, Prelude.Prelude__CREATE_Us_implHash_Sub_Hash_PreludeDot_NumHash_PreludeDot_FloatHash_( root.NewNode(  ), x5, Prelude.Prelude__CREATE_Us_implHash_Mul_Hash_PreludeDot_NumHash_PreludeDot_FloatHash_( root.NewNode(  ), x1, x4 ) ), Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) )
            return
        case 1:
            gocurry.ExemptCreate( root )
            return
    }
}

func Learn_adjustLayer( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    var x3 *gocurry.Node
    var x4 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    x3 = root.Children[ 2 ]
    x4 = root.Children[ 3 ]
    switch x2.GetConstructor(  ){
        case -1:
            if( task.IsBound( x2 ) ){
                task.ToHnf( x2 )
                return
            }else {
            }
            x2.SetTrLock( task.GetId(  ), gocurry.ChoiceCreate( task.NewNode(  ), Prelude.Prelude__CREATE_LSb_RSb_( task.NewNode(  ) ), Prelude.Prelude__CREATE_Col_( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) ) )
            return
        case 0:
            gocurry.ExemptCreate( root )
            return
        case 1:
            var x5 *gocurry.Node
            var x6 *gocurry.Node
            x5 = x2.Children[ 0 ]
            x6 = x2.Children[ 1 ]
            Learn__CREATE_adjustLayerUs_CASE0( root, x4, x6, x1, x5, x3 )
            return
    }
}

func Learn_adjustLayerUs_CASE0( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x4 *gocurry.Node
    var x6 *gocurry.Node
    var x1 *gocurry.Node
    var x5 *gocurry.Node
    var x3 *gocurry.Node
    x4 = root.Children[ 0 ]
    x6 = root.Children[ 1 ]
    x1 = root.Children[ 2 ]
    x5 = root.Children[ 3 ]
    x3 = root.Children[ 4 ]
    switch x4.GetConstructor(  ){
        case -1:
            if( task.IsBound( x4 ) ){
                task.ToHnf( x4 )
                return
            }else {
            }
            x4.SetTrLock( task.GetId(  ), gocurry.ChoiceCreate( task.NewNode(  ), Prelude.Prelude__CREATE_LSb_RSb_( task.NewNode(  ) ), Prelude.Prelude__CREATE_Col_( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) ) )
            return
        case 0:
            gocurry.ExemptCreate( root )
            return
        case 1:
            var x7 *gocurry.Node
            var x8 *gocurry.Node
            x7 = x4.Children[ 0 ]
            x8 = x4.Children[ 1 ]
            Learn__CREATE_adjustLayerUs_LET1( root, x6, x8, x1, x5, x3, x7 )
            return
    }
}

func Learn_adjustLayerUs_LET1( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x6 *gocurry.Node
    var x8 *gocurry.Node
    var x1 *gocurry.Node
    var x5 *gocurry.Node
    var x3 *gocurry.Node
    var x7 *gocurry.Node
    var x9 *gocurry.Node
    x6 = root.Children[ 0 ]
    x8 = root.Children[ 1 ]
    x1 = root.Children[ 2 ]
    x5 = root.Children[ 3 ]
    x3 = root.Children[ 4 ]
    x7 = root.Children[ 5 ]
    x9 = Learn__CREATE_adjustNeuron( root.NewNode(  ), x1, x5, x3, x7 )
    Learn__CREATE_adjustLayerUs_COMPLEXCASE2( root, x1, x6, x3, x8, x9, Prelude.Prelude__CREATE_Us_implHash_Eq_Eq_Hash_PreludeDot_EqHash_LSb_RSb_Hash_0Hash_Hash_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_instHash_PreludeDot_EqHash_LSb_RSb_Hash_0Hash_Hash_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_instHash_PreludeDot_EqHash_PreludeDot_FloatHash_( root.NewNode(  ) ) ), x6, Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) )
    return
}

func Learn_adjustLayerUs_COMPLEXCASE2( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x6 *gocurry.Node
    var x3 *gocurry.Node
    var x8 *gocurry.Node
    var x9 *gocurry.Node
    var x10 *gocurry.Node
    x1 = root.Children[ 0 ]
    x6 = root.Children[ 1 ]
    x3 = root.Children[ 2 ]
    x8 = root.Children[ 3 ]
    x9 = root.Children[ 4 ]
    x10 = root.Children[ 5 ]
    switch x10.GetConstructor(  ){
        case -1:
            if( task.IsBound( x10 ) ){
                task.ToHnf( x10 )
                return
            }else {
            }
            x10.SetTrLock( task.GetId(  ), gocurry.ChoiceCreate( task.NewNode(  ), Prelude.Prelude__CREATE_False( task.NewNode(  ) ), Prelude.Prelude__CREATE_True( task.NewNode(  ) ) ) )
            return
        case 0:
            Prelude.Prelude__CREATE_Col_( root, x9, Learn__CREATE_adjustLayer( root.NewNode(  ), x1, x6, x3, x8 ) )
            return
        case 1:
            Prelude.Prelude__CREATE_Col_( root, x9, Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) )
            return
    }
}

func Learn_gradientDescent( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    var x3 *gocurry.Node
    var x4 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    x3 = root.Children[ 2 ]
    x4 = root.Children[ 3 ]
    switch x2.GetConstructor(  ){
        case -1:
            if( task.IsBound( x2 ) ){
                task.ToHnf( x2 )
                return
            }else {
            }
            x2.SetTrLock( task.GetId(  ), gocurry.ChoiceCreate( task.NewNode(  ), Prelude.Prelude__CREATE_LSb_RSb_( task.NewNode(  ) ), Prelude.Prelude__CREATE_Col_( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) ) )
            return
        case 0:
            gocurry.ExemptCreate( root )
            return
        case 1:
            var x5 *gocurry.Node
            var x6 *gocurry.Node
            x5 = x2.Children[ 0 ]
            x6 = x2.Children[ 1 ]
            Learn__CREATE_gradientDescentUs_CASE0( root, x3, x4, x6, x1, x5 )
            return
    }
}

func Learn_gradientDescentUs_CASE0( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x3 *gocurry.Node
    var x4 *gocurry.Node
    var x6 *gocurry.Node
    var x1 *gocurry.Node
    var x5 *gocurry.Node
    x3 = root.Children[ 0 ]
    x4 = root.Children[ 1 ]
    x6 = root.Children[ 2 ]
    x1 = root.Children[ 3 ]
    x5 = root.Children[ 4 ]
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
            var x8 *gocurry.Node
            x8 = x3.Children[ 1 ]
            Learn__CREATE_gradientDescentUs_CASE1( root, x8, x4, x6, x1, x5 )
            return
    }
}

func Learn_gradientDescentUs_CASE1( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x8 *gocurry.Node
    var x4 *gocurry.Node
    var x6 *gocurry.Node
    var x1 *gocurry.Node
    var x5 *gocurry.Node
    x8 = root.Children[ 0 ]
    x4 = root.Children[ 1 ]
    x6 = root.Children[ 2 ]
    x1 = root.Children[ 3 ]
    x5 = root.Children[ 4 ]
    switch x8.GetConstructor(  ){
        case -1:
            if( task.IsBound( x8 ) ){
                task.ToHnf( x8 )
                return
            }else {
            }
            x8.SetTrLock( task.GetId(  ), gocurry.ChoiceCreate( task.NewNode(  ), Prelude.Prelude__CREATE_LSb_RSb_( task.NewNode(  ) ), Prelude.Prelude__CREATE_Col_( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) ) )
            return
        case 0:
            gocurry.ExemptCreate( root )
            return
        case 1:
            var x9 *gocurry.Node
            var x10 *gocurry.Node
            x9 = x8.Children[ 0 ]
            x10 = x8.Children[ 1 ]
            Learn__CREATE_gradientDescentUs_CASE2( root, x4, x6, x10, x1, x5, x9 )
            return
    }
}

func Learn_gradientDescentUs_CASE2( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x4 *gocurry.Node
    var x6 *gocurry.Node
    var x10 *gocurry.Node
    var x1 *gocurry.Node
    var x5 *gocurry.Node
    var x9 *gocurry.Node
    x4 = root.Children[ 0 ]
    x6 = root.Children[ 1 ]
    x10 = root.Children[ 2 ]
    x1 = root.Children[ 3 ]
    x5 = root.Children[ 4 ]
    x9 = root.Children[ 5 ]
    switch x4.GetConstructor(  ){
        case -1:
            if( task.IsBound( x4 ) ){
                task.ToHnf( x4 )
                return
            }else {
            }
            x4.SetTrLock( task.GetId(  ), gocurry.ChoiceCreate( task.NewNode(  ), Prelude.Prelude__CREATE_LSb_RSb_( task.NewNode(  ) ), Prelude.Prelude__CREATE_Col_( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) ) )
            return
        case 0:
            gocurry.ExemptCreate( root )
            return
        case 1:
            var x11 *gocurry.Node
            var x12 *gocurry.Node
            x11 = x4.Children[ 0 ]
            x12 = x4.Children[ 1 ]
            Learn__CREATE_gradientDescentUs_LET3( root, x6, x10, x12, x1, x5, x9, x11 )
            return
    }
}

func Learn_gradientDescentUs_LET3( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x6 *gocurry.Node
    var x10 *gocurry.Node
    var x12 *gocurry.Node
    var x1 *gocurry.Node
    var x5 *gocurry.Node
    var x9 *gocurry.Node
    var x11 *gocurry.Node
    var x13 *gocurry.Node
    x6 = root.Children[ 0 ]
    x10 = root.Children[ 1 ]
    x12 = root.Children[ 2 ]
    x1 = root.Children[ 3 ]
    x5 = root.Children[ 4 ]
    x9 = root.Children[ 5 ]
    x11 = root.Children[ 6 ]
    x13 = Learn__CREATE_adjustLayer( root.NewNode(  ), x1, x5, x9, x11 )
    Learn__CREATE_gradientDescentUs_COMPLEXCASE4( root, x1, x6, x9, x10, x12, x13, Prelude.Prelude__CREATE_Us_implHash_Eq_Eq_Hash_PreludeDot_EqHash_LSb_RSb_Hash_0Hash_Hash_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_instHash_PreludeDot_EqHash_LSb_RSb_Hash_0Hash_Hash_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_instHash_PreludeDot_EqHash_LSb_RSb_Hash_0Hash_Hash_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_instHash_PreludeDot_EqHash_PreludeDot_FloatHash_( root.NewNode(  ) ) ) ), x6, Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) ) )
    return
}

func Learn_gradientDescentUs_COMPLEXCASE4( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x6 *gocurry.Node
    var x9 *gocurry.Node
    var x10 *gocurry.Node
    var x12 *gocurry.Node
    var x13 *gocurry.Node
    var x14 *gocurry.Node
    x1 = root.Children[ 0 ]
    x6 = root.Children[ 1 ]
    x9 = root.Children[ 2 ]
    x10 = root.Children[ 3 ]
    x12 = root.Children[ 4 ]
    x13 = root.Children[ 5 ]
    x14 = root.Children[ 6 ]
    switch x14.GetConstructor(  ){
        case -1:
            if( task.IsBound( x14 ) ){
                task.ToHnf( x14 )
                return
            }else {
            }
            x14.SetTrLock( task.GetId(  ), gocurry.ChoiceCreate( task.NewNode(  ), Prelude.Prelude__CREATE_False( task.NewNode(  ) ), Prelude.Prelude__CREATE_True( task.NewNode(  ) ) ) )
            return
        case 0:
            Prelude.Prelude__CREATE_Col_( root, x13, Learn__CREATE_gradientDescent( root.NewNode(  ), x1, x6, Prelude.Prelude__CREATE_Col_( root.NewNode(  ), x9, x10 ), x12 ) )
            return
        case 1:
            Prelude.Prelude__CREATE_Col_( root, x13, Prelude.Prelude__CREATE_LSb_RSb_( root.NewNode(  ) ) )
            return
    }
}

func Learn_learnExample( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    var x3 *gocurry.Node
    var x4 *gocurry.Node
    var x5 *gocurry.Node
    var x6 *gocurry.Node
    var x7 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    x3 = root.Children[ 2 ]
    x4 = Perceptron.Perceptron__CREATE_feedForward( root.NewNode(  ), x1, Network.Network__CREATE_inputs( root.NewNode(  ), x3 ), x2 )
    x5 = Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_reverse( root.NewNode(  ) ), Network.Network__CREATE_layers( root.NewNode(  ), x1 ) )
    x6 = Network.Network__CREATE_learningRate( root.NewNode(  ), x1 )
    x7 = Learn__CREATE_backPropagate( root.NewNode(  ), x5, Network.Network__CREATE_expectation( root.NewNode(  ), x3 ), x2, x4 )
    Learn__CREATE_gradientDescent( root, x6, x2, x4, x7 )
    return
}

func Learn_learnExampleNetwork( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    Prelude.Prelude__CREATE_apply( root, x1, x2 )
    return
}

func Learn_learnEpic( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    var x3 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    x3 = Learn__CREATE_learnExampleNetwork( root.NewNode(  ), Learn__CREATE_learnExample( root.NewNode(  ) ), x1 )
    Prelude.Prelude__CREATE_foldl( root, x3, x2, Network.Network__CREATE_exs( root.NewNode(  ), x1 ) )
    return
}

func Learn_learn( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    var x3 *gocurry.Node
    var x4 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    x3 = root.Children[ 2 ]
    x4 = Learn__CREATE_learnEpic( root.NewNode(  ), x2, x3 )
    Learn__CREATE_learnUs_COMPLEXCASE0( root, x1, x2, x4, x3, Prelude.Prelude__CREATE_Us_implHash_Lt_Eq_Hash_PreludeDot_OrdHash_PreludeDot_IntHash_( root.NewNode(  ), x1, gocurry.IntLitCreate( root.NewNode(  ), 0 ) ) )
    return
}

func Learn_learnUs_COMPLEXCASE0( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    var x4 *gocurry.Node
    var x3 *gocurry.Node
    var x5 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    x4 = root.Children[ 2 ]
    x3 = root.Children[ 3 ]
    x5 = root.Children[ 4 ]
    switch x5.GetConstructor(  ){
        case -1:
            if( task.IsBound( x5 ) ){
                task.ToHnf( x5 )
                return
            }else {
            }
            x5.SetTrLock( task.GetId(  ), gocurry.ChoiceCreate( task.NewNode(  ), Prelude.Prelude__CREATE_False( task.NewNode(  ) ), Prelude.Prelude__CREATE_True( task.NewNode(  ) ) ) )
            return
        case 0:
            Learn__CREATE_learn( root, Prelude.Prelude__CREATE_Us_implHash_Sub_Hash_PreludeDot_NumHash_PreludeDot_IntHash_( root.NewNode(  ), x1, gocurry.IntLitCreate( root.NewNode(  ), 1 ) ), x2, x4 )
            return
        case 1:
            gocurry.RedirectCreate( root, x3 )
            return
    }
}

func Learn_initialWeights( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    Learn__CREATE_initialWeightsDot_initLWsDot_64( root, x2, Network.Network__CREATE_layers( root.NewNode(  ), x1 ) )
    return
}

func Learn_initialWeightsDot_initLWsDot_64( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    switch x2.GetConstructor(  ){
        case -1:
            if( task.IsBound( x2 ) ){
                task.ToHnf( x2 )
                return
            }else {
            }
            x2.SetTrLock( task.GetId(  ), gocurry.ChoiceCreate( task.NewNode(  ), Prelude.Prelude__CREATE_LSb_RSb_( task.NewNode(  ) ), Prelude.Prelude__CREATE_Col_( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) ) )
            return
        case 0:
            Prelude.Prelude__CREATE_LSb_RSb_( root )
            return
        case 1:
            var x3 *gocurry.Node
            var x4 *gocurry.Node
            x3 = x2.Children[ 0 ]
            x4 = x2.Children[ 1 ]
            Learn__CREATE_initialWeightsDot_initLWsDot_64Us_LET0( root, x1, x4, x3 )
            return
    }
}

func Learn_initialWeightsDot_initLWsDot_64Us_LET0( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x4 *gocurry.Node
    var x3 *gocurry.Node
    var x5 *gocurry.Node
    var x6 *gocurry.Node
    x1 = root.Children[ 0 ]
    x4 = root.Children[ 1 ]
    x3 = root.Children[ 2 ]
    x5 = Algorithms.Algorithms__CREATE_initWeights( root.NewNode(  ), Network.Network__CREATE_algorithm( root.NewNode(  ), x3 ) )
    x6 = Network.Network__CREATE_nodes( root.NewNode(  ), x3 )
    Prelude.Prelude__CREATE_Col_( root, Learn__CREATE_initialWeightsDot_initLWsDot_64Dot_initNWsDot_69( root.NewNode(  ), x3, x5, Prelude.Prelude__CREATE_Us_instHash_PreludeDot_NumHash_PreludeDot_IntHash_( root.NewNode(  ) ), Prelude.Prelude__CREATE_Us_instHash_PreludeDot_OrdHash_PreludeDot_IntHash_( root.NewNode(  ) ), x1, x6 ), Learn__CREATE_initialWeightsDot_initLWsDot_64( root.NewNode(  ), x6, x4 ) )
    return
}

func Learn_initialWeightsDot_initLWsDot_64Dot_initNWsDot_69( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    var x3 *gocurry.Node
    var x4 *gocurry.Node
    var x5 *gocurry.Node
    var x6 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    x3 = root.Children[ 2 ]
    x4 = root.Children[ 3 ]
    x5 = root.Children[ 4 ]
    x6 = root.Children[ 5 ]
    Learn__CREATE_initialWeightsDot_initLWsDot_64Dot_initNWsDot_69Us_COMPLEXCASE0( root, x1, x3, x4, x2, x5, x6, Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_Lt_Eq_( root.NewNode(  ), x4 ), x6 ), Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_fromInt( root.NewNode(  ), x3 ), gocurry.IntLitCreate( root.NewNode(  ), 0 ) ) ) )
    return
}

func Learn_initialWeightsDot_initLWsDot_64Dot_initNWsDot_69Us_COMPLEXCASE0( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x3 *gocurry.Node
    var x4 *gocurry.Node
    var x2 *gocurry.Node
    var x5 *gocurry.Node
    var x6 *gocurry.Node
    var x7 *gocurry.Node
    x1 = root.Children[ 0 ]
    x3 = root.Children[ 1 ]
    x4 = root.Children[ 2 ]
    x2 = root.Children[ 3 ]
    x5 = root.Children[ 4 ]
    x6 = root.Children[ 5 ]
    x7 = root.Children[ 6 ]
    switch x7.GetConstructor(  ){
        case -1:
            if( task.IsBound( x7 ) ){
                task.ToHnf( x7 )
                return
            }else {
            }
            x7.SetTrLock( task.GetId(  ), gocurry.ChoiceCreate( task.NewNode(  ), Prelude.Prelude__CREATE_False( task.NewNode(  ) ), Prelude.Prelude__CREATE_True( task.NewNode(  ) ) ) )
            return
        case 0:
            Learn__CREATE_initialWeightsDot_initLWsDot_64Dot_initNWsDot_69Dot_initNWsSQuote_Dot_72( root, x1, x3, x4, x2, x5, x6 )
            return
        case 1:
            Prelude.Prelude__CREATE_LSb_RSb_( root )
            return
    }
}

func Learn_initialWeightsDot_initLWsDot_64Dot_initNWsDot_69Dot_initNWsSQuote_Dot_72( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    var x3 *gocurry.Node
    var x4 *gocurry.Node
    var x5 *gocurry.Node
    var x6 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    x3 = root.Children[ 2 ]
    x4 = root.Children[ 3 ]
    x5 = root.Children[ 4 ]
    x6 = root.Children[ 5 ]
    Prelude.Prelude__CREATE_Col_( root, Learn__CREATE_initialWeightsDot_initLWsDot_64Dot_initNWsDot_69Dot_initNWsSQuote_Dot_72Dot_initWsDot_74( root.NewNode(  ), x1, x4, x5 ), Learn__CREATE_initialWeightsDot_initLWsDot_64Dot_initNWsDot_69( root.NewNode(  ), x1, x4, x2, x3, x5, Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_Sub_( root.NewNode(  ), x2 ), x6 ), Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_fromInt( root.NewNode(  ), x2 ), gocurry.IntLitCreate( root.NewNode(  ), 1 ) ) ) ) )
    return
}

func Learn_initialWeightsDot_initLWsDot_64Dot_initNWsDot_69Dot_initNWsSQuote_Dot_72Dot_initWsDot_74( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    var x3 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    x3 = root.Children[ 2 ]
    Prelude.Prelude__CREATE_replicate( root, Prelude.Prelude__CREATE_Us_implHash_Add_Hash_PreludeDot_NumHash_PreludeDot_IntHash_( root.NewNode(  ), x3, gocurry.IntLitCreate( root.NewNode(  ), 1 ) ), Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_apply( root.NewNode(  ), x2, x3 ), Network.Network__CREATE_nodes( root.NewNode(  ), x1 ) ) )
    return
}

