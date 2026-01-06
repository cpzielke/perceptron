package Perceptron

import "gocurry"
import "curry2go/Algorithms"
import "curry2go/Network"
import "curry2go/Prelude"

var func_names []string = []string{ "dotProduct", "dotProduct_CASE2", "dotProduct_CASE3", "dotProduct_CASE0", "dotProduct_CASE1", "layerDotProduct", "networkDotProduct", "networkDotProduct.accum.24", "networkDotProduct.accum.24_CASE0", "networkDotProduct.accum.24_LET1", "feedForward", "eval", "eval._#selFP2#out", "eval2", "eval2._#selFP4#out" }

func Perceptron__CREATE_dotProduct( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Perceptron_dotProduct, & func_names[ 0 ], 2, 0, args... )
    return( root )
}

func Perceptron__CREATE_dotProductUs_CASE2( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Perceptron_dotProductUs_CASE2, & func_names[ 1 ], 3, 0, args... )
    return( root )
}

func Perceptron__CREATE_dotProductUs_CASE3( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Perceptron_dotProductUs_CASE3, & func_names[ 2 ], 1, 0, args... )
    return( root )
}

func Perceptron__CREATE_dotProductUs_CASE0( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Perceptron_dotProductUs_CASE0, & func_names[ 3 ], 1, 0, args... )
    return( root )
}

func Perceptron__CREATE_dotProductUs_CASE1( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Perceptron_dotProductUs_CASE1, & func_names[ 4 ], 2, 0, args... )
    return( root )
}

func Perceptron__CREATE_layerDotProduct( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Perceptron_layerDotProduct, & func_names[ 5 ], 3, 2, args... )
    return( root )
}

func Perceptron__CREATE_networkDotProduct( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Perceptron_networkDotProduct, & func_names[ 6 ], 3, -1, args... )
    return( root )
}

func Perceptron__CREATE_networkDotProductDot_accumDot_24( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Perceptron_networkDotProductDot_accumDot_24, & func_names[ 7 ], 3, 0, args... )
    return( root )
}

func Perceptron__CREATE_networkDotProductDot_accumDot_24Us_CASE0( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Perceptron_networkDotProductDot_accumDot_24Us_CASE0, & func_names[ 8 ], 4, 0, args... )
    return( root )
}

func Perceptron__CREATE_networkDotProductDot_accumDot_24Us_LET1( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Perceptron_networkDotProductDot_accumDot_24Us_LET1, & func_names[ 9 ], 5, -1, args... )
    return( root )
}

func Perceptron__CREATE_feedForward( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Perceptron_feedForward, & func_names[ 10 ], 3, -1, args... )
    return( root )
}

func Perceptron__CREATE_eval( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Perceptron_eval, & func_names[ 11 ], 2, -1, args... )
    return( root )
}

func Perceptron__CREATE_evalDot_Us_Hash_selFP2Hash_out( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Perceptron_evalDot_Us_Hash_selFP2Hash_out, & func_names[ 12 ], 1, 0, args... )
    return( root )
}

func Perceptron__CREATE_eval2( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Perceptron_eval2, & func_names[ 13 ], 3, -1, args... )
    return( root )
}

func Perceptron__CREATE_eval2Dot_Us_Hash_selFP4Hash_out( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Perceptron_eval2Dot_Us_Hash_selFP4Hash_out, & func_names[ 14 ], 1, 0, args... )
    return( root )
}

func Perceptron_dotProduct( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
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
            Perceptron__CREATE_dotProductUs_CASE0( root, x2 )
            return
        case 1:
            var x3 *gocurry.Node
            var x4 *gocurry.Node
            x3 = x1.Children[ 0 ]
            x4 = x1.Children[ 1 ]
            Perceptron__CREATE_dotProductUs_CASE2( root, x2, x3, x4 )
            return
    }
}

func Perceptron_dotProductUs_CASE2( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x2 *gocurry.Node
    var x3 *gocurry.Node
    var x4 *gocurry.Node
    x2 = root.Children[ 0 ]
    x3 = root.Children[ 1 ]
    x4 = root.Children[ 2 ]
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
            Perceptron__CREATE_dotProductUs_CASE3( root, x4 )
            return
        case 1:
            var x7 *gocurry.Node
            var x8 *gocurry.Node
            x7 = x2.Children[ 0 ]
            x8 = x2.Children[ 1 ]
            Prelude.Prelude__CREATE_Us_implHash_Add_Hash_PreludeDot_NumHash_PreludeDot_FloatHash_( root, Prelude.Prelude__CREATE_Us_implHash_Mul_Hash_PreludeDot_NumHash_PreludeDot_FloatHash_( root.NewNode(  ), x3, x7 ), Perceptron__CREATE_dotProduct( root.NewNode(  ), x4, x8 ) )
            return
    }
}

func Perceptron_dotProductUs_CASE3( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x4 *gocurry.Node
    x4 = root.Children[ 0 ]
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
            gocurry.ExemptCreate( root )
            return
    }
}

func Perceptron_dotProductUs_CASE0( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x2 *gocurry.Node
    x2 = root.Children[ 0 ]
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
            gocurry.FloatLitCreate( root, 0 )
            return
        case 1:
            var x9 *gocurry.Node
            var x10 *gocurry.Node
            x9 = x2.Children[ 0 ]
            x10 = x2.Children[ 1 ]
            Perceptron__CREATE_dotProductUs_CASE1( root, x10, x9 )
            return
    }
}

func Perceptron_dotProductUs_CASE1( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x10 *gocurry.Node
    var x9 *gocurry.Node
    x10 = root.Children[ 0 ]
    x9 = root.Children[ 1 ]
    switch x10.GetConstructor(  ){
        case -1:
            if( task.IsBound( x10 ) ){
                task.ToHnf( x10 )
                return
            }else {
            }
            x10.SetTrLock( task.GetId(  ), gocurry.ChoiceCreate( task.NewNode(  ), Prelude.Prelude__CREATE_LSb_RSb_( task.NewNode(  ) ), Prelude.Prelude__CREATE_Col_( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) ) )
            return
        case 0:
            gocurry.RedirectCreate( root, x9 )
            return
        case 1:
            gocurry.ExemptCreate( root )
            return
    }
}

func Perceptron_layerDotProduct( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    var x3 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    x3 = root.Children[ 2 ]
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
            var x4 *gocurry.Node
            var x5 *gocurry.Node
            x4 = x3.Children[ 0 ]
            x5 = x3.Children[ 1 ]
            Prelude.Prelude__CREATE_Col_( root, Prelude.Prelude__CREATE_apply( root.NewNode(  ), x1, Perceptron__CREATE_dotProduct( root.NewNode(  ), x2, x4 ) ), Perceptron__CREATE_layerDotProduct( root.NewNode(  ), x1, x2, x5 ) )
            return
    }
}

func Perceptron_networkDotProduct( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    var x3 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    x3 = root.Children[ 2 ]
    Prelude.Prelude__CREATE_Col_( root, x2, Perceptron__CREATE_networkDotProductDot_accumDot_24( root.NewNode(  ), Network.Network__CREATE_layers( root.NewNode(  ), x1 ), x2, x3 ) )
    return
}

func Perceptron_networkDotProductDot_accumDot_24( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    var x3 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    x3 = root.Children[ 2 ]
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
            Prelude.Prelude__CREATE_LSb_RSb_( root )
            return
        case 1:
            var x4 *gocurry.Node
            var x5 *gocurry.Node
            x4 = x1.Children[ 0 ]
            x5 = x1.Children[ 1 ]
            Perceptron__CREATE_networkDotProductDot_accumDot_24Us_CASE0( root, x3, x5, x4, x2 )
            return
    }
}

func Perceptron_networkDotProductDot_accumDot_24Us_CASE0( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x3 *gocurry.Node
    var x5 *gocurry.Node
    var x4 *gocurry.Node
    var x2 *gocurry.Node
    x3 = root.Children[ 0 ]
    x5 = root.Children[ 1 ]
    x4 = root.Children[ 2 ]
    x2 = root.Children[ 3 ]
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
            var x6 *gocurry.Node
            var x7 *gocurry.Node
            x6 = x3.Children[ 0 ]
            x7 = x3.Children[ 1 ]
            Perceptron__CREATE_networkDotProductDot_accumDot_24Us_LET1( root, x5, x7, x4, x2, x6 )
            return
    }
}

func Perceptron_networkDotProductDot_accumDot_24Us_LET1( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x5 *gocurry.Node
    var x7 *gocurry.Node
    var x4 *gocurry.Node
    var x2 *gocurry.Node
    var x6 *gocurry.Node
    var x8 *gocurry.Node
    var x9 *gocurry.Node
    x5 = root.Children[ 0 ]
    x7 = root.Children[ 1 ]
    x4 = root.Children[ 2 ]
    x2 = root.Children[ 3 ]
    x6 = root.Children[ 4 ]
    x8 = Algorithms.Algorithms__CREATE_activation( root.NewNode(  ), Network.Network__CREATE_algorithm( root.NewNode(  ), x4 ) )
    x9 = Perceptron__CREATE_layerDotProduct( root.NewNode(  ), x8, x2, x6 )
    Prelude.Prelude__CREATE_Col_( root, x9, Perceptron__CREATE_networkDotProductDot_accumDot_24( root.NewNode(  ), x5, x9, x7 ) )
    return
}

func Perceptron_feedForward( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    var x3 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    x3 = root.Children[ 2 ]
    Prelude.Prelude__CREATE_apply( root, Prelude.Prelude__CREATE_reverse( root.NewNode(  ) ), Perceptron__CREATE_networkDotProduct( root.NewNode(  ), x1, x2, Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_reverse( root.NewNode(  ) ), x3 ) ) )
    return
}

func Perceptron_eval( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    var x3 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    x3 = Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_reverse( root.NewNode(  ) ), Perceptron__CREATE_networkDotProduct( root.NewNode(  ), x1, x2, Network.Network__CREATE_weights( root.NewNode(  ), x1 ) ) )
    Perceptron__CREATE_evalDot_Us_Hash_selFP2Hash_out( root, x3 )
    return
}

func Perceptron_evalDot_Us_Hash_selFP2Hash_out( task *gocurry.Task )(  ){
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
            var x2 *gocurry.Node
            x2 = x1.Children[ 0 ]
            gocurry.RedirectCreate( root, x2 )
            return
    }
}

func Perceptron_eval2( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    var x3 *gocurry.Node
    var x4 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    x3 = root.Children[ 2 ]
    x4 = Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_reverse( root.NewNode(  ) ), Perceptron__CREATE_networkDotProduct( root.NewNode(  ), x1, x3, x2 ) )
    Perceptron__CREATE_eval2Dot_Us_Hash_selFP4Hash_out( root, x4 )
    return
}

func Perceptron_eval2Dot_Us_Hash_selFP4Hash_out( task *gocurry.Task )(  ){
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
            var x2 *gocurry.Node
            x2 = x1.Children[ 0 ]
            gocurry.RedirectCreate( root, x2 )
            return
    }
}

