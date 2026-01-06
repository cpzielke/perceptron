package Network

import "gocurry"
import "curry2go/Prelude"

var func_names []string = []string{ "_inst#Prelude.Data#Network.Example#", "_impl#===#Prelude.Data#Network.Example#", "_impl#===#Prelude.Data#Network.Example#_CASE0", "_impl#aValue#Prelude.Data#Network.Example#", "name", "inputs", "expectation", "nodes", "algorithm", "weights", "exs", "layers", "learningRate" }

var Network_Example_names []string = []string{ "Example" }

var Network_Layer_names []string = []string{ "Layer" }

var Network_Network_names []string = []string{ "Network" }

func Network__CREATE_Example( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.ConstCreate( root, 0, 3, & Network_Example_names[ 0 ], args... )
    return( root )
}

func Network__CREATE_Layer( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.ConstCreate( root, 0, 2, & Network_Layer_names[ 0 ], args... )
    return( root )
}

func Network__CREATE_Network( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.ConstCreate( root, 0, 4, & Network_Network_names[ 0 ], args... )
    return( root )
}

func Network__CREATE_Us_instHash_PreludeDot_DataHash_NetworkDot_ExampleHash_( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Network_Us_instHash_PreludeDot_DataHash_NetworkDot_ExampleHash_, & func_names[ 0 ], 1, 0, args... )
    return( root )
}

func Network__CREATE_Us_implHash_Eq_Eq_Eq_Hash_PreludeDot_DataHash_NetworkDot_ExampleHash_( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Network_Us_implHash_Eq_Eq_Eq_Hash_PreludeDot_DataHash_NetworkDot_ExampleHash_, & func_names[ 1 ], 2, 0, args... )
    return( root )
}

func Network__CREATE_Us_implHash_Eq_Eq_Eq_Hash_PreludeDot_DataHash_NetworkDot_ExampleHash_Us_CASE0( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Network_Us_implHash_Eq_Eq_Eq_Hash_PreludeDot_DataHash_NetworkDot_ExampleHash_Us_CASE0, & func_names[ 2 ], 4, 0, args... )
    return( root )
}

func Network__CREATE_Us_implHash_aValueHash_PreludeDot_DataHash_NetworkDot_ExampleHash_( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Network_Us_implHash_aValueHash_PreludeDot_DataHash_NetworkDot_ExampleHash_, & func_names[ 3 ], 0, -1, args... )
    return( root )
}

func Network__CREATE_name( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Network_name, & func_names[ 4 ], 1, 0, args... )
    return( root )
}

func Network__CREATE_inputs( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Network_inputs, & func_names[ 5 ], 1, 0, args... )
    return( root )
}

func Network__CREATE_expectation( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Network_expectation, & func_names[ 6 ], 1, 0, args... )
    return( root )
}

func Network__CREATE_nodes( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Network_nodes, & func_names[ 7 ], 1, 0, args... )
    return( root )
}

func Network__CREATE_algorithm( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Network_algorithm, & func_names[ 8 ], 1, 0, args... )
    return( root )
}

func Network__CREATE_weights( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Network_weights, & func_names[ 9 ], 1, 0, args... )
    return( root )
}

func Network__CREATE_exs( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Network_exs, & func_names[ 10 ], 1, 0, args... )
    return( root )
}

func Network__CREATE_layers( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Network_layers, & func_names[ 11 ], 1, 0, args... )
    return( root )
}

func Network__CREATE_learningRate( root *gocurry.Node, args ...*gocurry.Node )( *gocurry.Node ){
    gocurry.FuncCreate( root, Network_learningRate, & func_names[ 12 ], 1, 0, args... )
    return( root )
}

func Network_Us_instHash_PreludeDot_DataHash_NetworkDot_ExampleHash_( task *gocurry.Task )(  ){
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
            x1.SetTrLock( task.GetId(  ), Prelude.Prelude__CREATE_Lb_Rb_( task.NewNode(  ) ) )
            return
        case 0:
            Prelude.Prelude__CREATE_Us_DictHash_Data( root, Network__CREATE_Us_implHash_Eq_Eq_Eq_Hash_PreludeDot_DataHash_NetworkDot_ExampleHash_( root.NewNode(  ) ), Network__CREATE_Us_implHash_aValueHash_PreludeDot_DataHash_NetworkDot_ExampleHash_( root.NewNode(  ) ) )
            return
    }
}

func Network_Us_implHash_Eq_Eq_Eq_Hash_PreludeDot_DataHash_NetworkDot_ExampleHash_( task *gocurry.Task )(  ){
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
            x1.SetTrLock( task.GetId(  ), Network__CREATE_Example( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) )
            return
        case 0:
            var x3 *gocurry.Node
            var x4 *gocurry.Node
            var x5 *gocurry.Node
            x3 = x1.Children[ 0 ]
            x4 = x1.Children[ 1 ]
            x5 = x1.Children[ 2 ]
            Network__CREATE_Us_implHash_Eq_Eq_Eq_Hash_PreludeDot_DataHash_NetworkDot_ExampleHash_Us_CASE0( root, x2, x3, x4, x5 )
            return
    }
}

func Network_Us_implHash_Eq_Eq_Eq_Hash_PreludeDot_DataHash_NetworkDot_ExampleHash_Us_CASE0( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    var x2 *gocurry.Node
    var x3 *gocurry.Node
    var x4 *gocurry.Node
    var x5 *gocurry.Node
    x2 = root.Children[ 0 ]
    x3 = root.Children[ 1 ]
    x4 = root.Children[ 2 ]
    x5 = root.Children[ 3 ]
    switch x2.GetConstructor(  ){
        case -1:
            if( task.IsBound( x2 ) ){
                task.ToHnf( x2 )
                return
            }else {
            }
            x2.SetTrLock( task.GetId(  ), Network__CREATE_Example( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) )
            return
        case 0:
            var x6 *gocurry.Node
            var x7 *gocurry.Node
            var x8 *gocurry.Node
            x6 = x2.Children[ 0 ]
            x7 = x2.Children[ 1 ]
            x8 = x2.Children[ 2 ]
            Prelude.Prelude__CREATE_And_And_( root, Prelude.Prelude__CREATE_And_And_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_implHash_Eq_Eq_Eq_Hash_PreludeDot_DataHash_LSb_RSb_Hash_0Hash_Hash_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_instHash_PreludeDot_DataHash_PreludeDot_CharHash_( root.NewNode(  ) ), x3, x6 ), Prelude.Prelude__CREATE_Us_implHash_Eq_Eq_Eq_Hash_PreludeDot_DataHash_LSb_RSb_Hash_0Hash_Hash_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_instHash_PreludeDot_DataHash_PreludeDot_FloatHash_( root.NewNode(  ) ), x4, x7 ) ), Prelude.Prelude__CREATE_Us_implHash_Eq_Eq_Eq_Hash_PreludeDot_DataHash_LSb_RSb_Hash_0Hash_Hash_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_instHash_PreludeDot_DataHash_PreludeDot_FloatHash_( root.NewNode(  ) ), x5, x8 ) )
            return
    }
}

func Network_Us_implHash_aValueHash_PreludeDot_DataHash_NetworkDot_ExampleHash_( task *gocurry.Task )(  ){
    root := task.GetControl(  )
    Network__CREATE_Example( root, Prelude.Prelude__CREATE_Us_implHash_aValueHash_PreludeDot_DataHash_LSb_RSb_Hash_0Hash_Hash_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_instHash_PreludeDot_DataHash_PreludeDot_CharHash_( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_Us_implHash_aValueHash_PreludeDot_DataHash_LSb_RSb_Hash_0Hash_Hash_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_instHash_PreludeDot_DataHash_PreludeDot_FloatHash_( root.NewNode(  ) ) ), Prelude.Prelude__CREATE_Us_implHash_aValueHash_PreludeDot_DataHash_LSb_RSb_Hash_0Hash_Hash_( root.NewNode(  ), Prelude.Prelude__CREATE_Us_instHash_PreludeDot_DataHash_PreludeDot_FloatHash_( root.NewNode(  ) ) ) )
    return
}

func Network_name( task *gocurry.Task )(  ){
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
            x1.SetTrLock( task.GetId(  ), Network__CREATE_Example( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) )
            return
        case 0:
            var x2 *gocurry.Node
            x2 = x1.Children[ 0 ]
            gocurry.RedirectCreate( root, x2 )
            return
    }
}

func Network_inputs( task *gocurry.Task )(  ){
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
            x1.SetTrLock( task.GetId(  ), Network__CREATE_Example( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) )
            return
        case 0:
            var x3 *gocurry.Node
            x3 = x1.Children[ 1 ]
            gocurry.RedirectCreate( root, x3 )
            return
    }
}

func Network_expectation( task *gocurry.Task )(  ){
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
            x1.SetTrLock( task.GetId(  ), Network__CREATE_Example( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) )
            return
        case 0:
            var x4 *gocurry.Node
            x4 = x1.Children[ 2 ]
            gocurry.RedirectCreate( root, x4 )
            return
    }
}

func Network_nodes( task *gocurry.Task )(  ){
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
            x1.SetTrLock( task.GetId(  ), Network__CREATE_Layer( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) )
            return
        case 0:
            var x2 *gocurry.Node
            x2 = x1.Children[ 0 ]
            gocurry.RedirectCreate( root, x2 )
            return
    }
}

func Network_algorithm( task *gocurry.Task )(  ){
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
            x1.SetTrLock( task.GetId(  ), Network__CREATE_Layer( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) )
            return
        case 0:
            var x3 *gocurry.Node
            x3 = x1.Children[ 1 ]
            gocurry.RedirectCreate( root, x3 )
            return
    }
}

func Network_weights( task *gocurry.Task )(  ){
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
            x1.SetTrLock( task.GetId(  ), Network__CREATE_Network( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) )
            return
        case 0:
            var x2 *gocurry.Node
            x2 = x1.Children[ 0 ]
            gocurry.RedirectCreate( root, x2 )
            return
    }
}

func Network_exs( task *gocurry.Task )(  ){
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
            x1.SetTrLock( task.GetId(  ), Network__CREATE_Network( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) )
            return
        case 0:
            var x3 *gocurry.Node
            x3 = x1.Children[ 1 ]
            gocurry.RedirectCreate( root, x3 )
            return
    }
}

func Network_layers( task *gocurry.Task )(  ){
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
            x1.SetTrLock( task.GetId(  ), Network__CREATE_Network( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) )
            return
        case 0:
            var x4 *gocurry.Node
            x4 = x1.Children[ 2 ]
            gocurry.RedirectCreate( root, x4 )
            return
    }
}

func Network_learningRate( task *gocurry.Task )(  ){
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
            x1.SetTrLock( task.GetId(  ), Network__CREATE_Network( task.NewNode(  ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ), gocurry.FreeCreate( task.NewNode(  ) ) ) )
            return
        case 0:
            var x5 *gocurry.Node
            x5 = x1.Children[ 3 ]
            gocurry.RedirectCreate( root, x5 )
            return
    }
}

