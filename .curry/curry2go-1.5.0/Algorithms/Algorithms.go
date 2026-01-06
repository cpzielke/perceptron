package Algorithms

import "gocurry"
import "curry2go/Prelude"

var func_names []string = []string{ "activation", "activation_der", "initWeights", "e", "classifier_act", "classifier_act_COMPLEXCASE0", "classifier_der", "classifier_iw", "classifier", "logistic_act", "logistic_der", "logistic_iw", "logistic", "scaling" }

var Algorithms_Algorithm_names []string = []string{ "Algorithm" }

func Algorithms__CREATE_Algorithm( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.ConstCreate( root, 0, 3, & Algorithms_Algorithm_names[ 0 ], args... )
    return( root )
}

func Algorithms__CREATE_activation( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Algorithms_activation, & func_names[ 0 ], 1, 0, args... )
    return( root )
}

func Algorithms__CREATE_activationUs_der( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Algorithms_activationUs_der, & func_names[ 1 ], 1, 0, args... )
    return( root )
}

func Algorithms__CREATE_initWeights( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Algorithms_initWeights, & func_names[ 2 ], 1, 0, args... )
    return( root )
}

func Algorithms__CREATE_e( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Algorithms_e, & func_names[ 3 ], 0, -1, args... )
    return( root )
}

func Algorithms__CREATE_classifierUs_act( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Algorithms_classifierUs_act, & func_names[ 4 ], 1, -1, args... )
    return( root )
}

func Algorithms__CREATE_classifierUs_actUs_COMPLEXCASE0( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Algorithms_classifierUs_actUs_COMPLEXCASE0, & func_names[ 5 ], 1, 0, args... )
    return( root )
}

func Algorithms__CREATE_classifierUs_der( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Algorithms_classifierUs_der, & func_names[ 6 ], 1, -1, args... )
    return( root )
}

func Algorithms__CREATE_classifierUs_iw( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Algorithms_classifierUs_iw, & func_names[ 7 ], 2, -1, args... )
    return( root )
}

func Algorithms__CREATE_classifier( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Algorithms_classifier, & func_names[ 8 ], 0, -1, args... )
    return( root )
}

func Algorithms__CREATE_logisticUs_act( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Algorithms_logisticUs_act, & func_names[ 9 ], 1, -1, args... )
    return( root )
}

func Algorithms__CREATE_logisticUs_der( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Algorithms_logisticUs_der, & func_names[ 10 ], 1, -1, args... )
    return( root )
}

func Algorithms__CREATE_logisticUs_iw( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Algorithms_logisticUs_iw, & func_names[ 11 ], 2, -1, args... )
    return( root )
}

func Algorithms__CREATE_logistic( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Algorithms_logistic, & func_names[ 12 ], 0, -1, args... )
    return( root )
}

func Algorithms__CREATE_scaling( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Algorithms_scaling, & func_names[ 13 ], 1, -1, args... )
    return( root )
}

func Algorithms_activation( task *gocurry.Task )(  ){
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
            x1.SetTrLock( task.GetId(  ), Algorithms__CREATE_Algorithm( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) )
            return
        case 0:
            var x2 *gocurry.Node
            x2 = x1.Children[ 0 ]
            gocurry.RedirectCreate( root, x2 )
            return
    }
}

func Algorithms_activationUs_der( task *gocurry.Task )(  ){
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
            x1.SetTrLock( task.GetId(  ), Algorithms__CREATE_Algorithm( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) )
            return
        case 0:
            var x3 *gocurry.Node
            x3 = x1.Children[ 1 ]
            gocurry.RedirectCreate( root, x3 )
            return
    }
}

func Algorithms_initWeights( task *gocurry.Task )(  ){
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
            x1.SetTrLock( task.GetId(  ), Algorithms__CREATE_Algorithm( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) )
            return
        case 0:
            var x4 *gocurry.Node
            x4 = x1.Children[ 2 ]
            gocurry.RedirectCreate( root, x4 )
            return
    }
}

func Algorithms_e( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    gocurry.FloatLitCreate( root, 2.718281828459045 )
    return
}

func Algorithms_classifierUs_act( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    x1 = root.Children[ 0 ]
    Algorithms__CREATE_classifierUs_actUs_COMPLEXCASE0( root, Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_Us_implHash_Gt_Eq_Hash_PreludeDot_OrdHash_PreludeDot_FloatHash_( root.NewNode(  ) ), x1 ), gocurry.FloatLitCreate( root.NewNode(  ), 0 ) ) )
    return
}

func Algorithms_classifierUs_actUs_COMPLEXCASE0( task *gocurry.Task )(  ){
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
            x2.SetTrLock( task.GetId(  ), gocurry.ChoiceCreate( task.NewNode(  ), Prelude.Prelude__CREATE_False( task.NewNode(  ) ), Prelude.Prelude__CREATE_True( task.NewNode(  ) ) ) )
            return
        case 0:
            gocurry.FloatLitCreate( root, 0 )
            return
        case 1:
            gocurry.FloatLitCreate( root, 1 )
            return
    }
}

func Algorithms_classifierUs_der( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    gocurry.FloatLitCreate( root, 1 )
    return
}

func Algorithms_classifierUs_iw( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    Algorithms__CREATE_scaling( root, Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_Us_implHash_sqrtHash_PreludeDot_FloatingHash_PreludeDot_FloatHash_( root.NewNode(  ) ), Prelude.Prelude__CREATE_Us_implHash_Slash_Hash_PreludeDot_FractionalHash_PreludeDot_FloatHash_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 7 ), Prelude.Prelude__CREATE_Us_implHash_Add_Hash_PreludeDot_NumHash_PreludeDot_FloatHash_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_implHash_toFloatHash_PreludeDot_RealHash_PreludeDot_IntHash_( root.NewNode(  ), x1 ), Prelude.Prelude__CREATE_Us_implHash_toFloatHash_PreludeDot_RealHash_PreludeDot_IntHash_( root.NewNode(  ), x2 ) ) ) ) )
    return
}

func Algorithms_classifier( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Algorithms__CREATE_Algorithm( root, Algorithms__CREATE_classifierUs_act( root.NewNode(  ) ), Algorithms__CREATE_classifierUs_der( root.NewNode(  ) ), Algorithms__CREATE_classifierUs_iw( root.NewNode(  ) ) )
    return
}

func Algorithms_logisticUs_act( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    x1 = root.Children[ 0 ]
    Prelude.Prelude__CREATE_Us_implHash_Slash_Hash_PreludeDot_FractionalHash_PreludeDot_FloatHash_( root, gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_Us_implHash_Add_Hash_PreludeDot_NumHash_PreludeDot_FloatHash_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ), Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_Us_implHash_Mul_Mul_Hash_PreludeDot_FloatingHash_PreludeDot_FloatHash_( root.NewNode(  ) ), Algorithms__CREATE_e( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_Us_implHash_negateHash_PreludeDot_NumHash_PreludeDot_FloatHash_( root.NewNode(  ), x1 ) ) ) )
    return
}

func Algorithms_logisticUs_der( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    x1 = root.Children[ 0 ]
    Prelude.Prelude__CREATE_Us_implHash_Sub_Hash_PreludeDot_NumHash_PreludeDot_FloatHash_( root, x1, Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_Us_implHash_Mul_Mul_Hash_PreludeDot_FloatingHash_PreludeDot_FloatHash_( root.NewNode(  ) ), x1 ), gocurry.FloatLitCreate( root.NewNode(  ), 2 ) ) )
    return
}

func Algorithms_logisticUs_iw( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    var x2 *gocurry.Node
    x1 = root.Children[ 0 ]
    x2 = root.Children[ 1 ]
    Algorithms__CREATE_scaling( root, Prelude.Prelude__CREATE_apply( root.NewNode(  ), Prelude.Prelude__CREATE_Us_implHash_sqrtHash_PreludeDot_FloatingHash_PreludeDot_FloatHash_( root.NewNode(  ) ), Prelude.Prelude__CREATE_Us_implHash_Slash_Hash_PreludeDot_FractionalHash_PreludeDot_FloatHash_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 6 ), Prelude.Prelude__CREATE_Us_implHash_Add_Hash_PreludeDot_NumHash_PreludeDot_FloatHash_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_implHash_toFloatHash_PreludeDot_RealHash_PreludeDot_IntHash_( root.NewNode(  ), x1 ), Prelude.Prelude__CREATE_Us_implHash_toFloatHash_PreludeDot_RealHash_PreludeDot_IntHash_( root.NewNode(  ), x2 ) ) ) ) )
    return
}

func Algorithms_logistic( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Algorithms__CREATE_Algorithm( root, Algorithms__CREATE_logisticUs_act( root.NewNode(  ) ), Algorithms__CREATE_logisticUs_der( root.NewNode(  ) ), Algorithms__CREATE_logisticUs_iw( root.NewNode(  ) ) )
    return
}

func Algorithms_scaling( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x1 *gocurry.Node
    x1 = root.Children[ 0 ]
    Prelude.Prelude__CREATE_Us_implHash_Slash_Hash_PreludeDot_FractionalHash_PreludeDot_FloatHash_( root, Prelude.Prelude__CREATE_Us_implHash_Sub_Hash_PreludeDot_NumHash_PreludeDot_FloatHash_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_implHash_Add_Hash_PreludeDot_NumHash_PreludeDot_FloatHash_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_implHash_Mul_Hash_PreludeDot_NumHash_PreludeDot_FloatHash_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_implHash_Mul_Hash_PreludeDot_NumHash_PreludeDot_FloatHash_( root.NewNode(  ), gocurry.FloatLitCreate( root.NewNode(  ), 2 ), x1 ), gocurry.FloatLitCreate( root.NewNode(  ), 10 ) ), gocurry.FloatLitCreate( root.NewNode(  ), 1 ) ), Prelude.Prelude__CREATE_Us_implHash_Mul_Hash_PreludeDot_NumHash_PreludeDot_FloatHash_( root.NewNode(  ), x1, gocurry.FloatLitCreate( root.NewNode(  ), 10 ) ) ), gocurry.FloatLitCreate( root.NewNode(  ), 50 ) )
    return
}

