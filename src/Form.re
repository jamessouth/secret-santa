// [@bs.val] external fetch: string => Js.Promise.t('a) = "fetch";




[@react.component]
let make = () => {
    // let (state, dispatch) = React.useReducer(reducer, initialState);
    let (inputValue, setInputValue) = React.useState(() => "");




    // let getData = (inp) => {
    //     Js.log3("looppoopol", inp, state);
    //     Js.Promise.(
    //         fetch("http://localhost:8000/?id=" ++ inp)
    //         |> then_(response => response##json())
    //         |> then_(jsonResponse => {
    //             Js.log2("re5sss", jsonResponse);
    //             dispatch(LoadedData(jsonResponse));
    //             Js.Promise.resolve();
    //             })
    //         |> catch(_err => {
    //             dispatch(ErrorFetchingData);
    //             Js.Promise.resolve();
    //             })
    //         |> ignore
    //     );
    // };

    <form action="http://localhost:8000" method="POST">
        <div>
            <input
                onChange={
                    event => {
                        let value = ReactEvent.Form.target(event)##value;
                        setInputValue(_ => value)
                    }
                }
                value=inputValue
                name="email"
            />
            <button
                // type_="button"
                // onClick={_ => {
                //     getData(inputValue)
                // }}
            >
                "Submit"->React.string
            </button>



            <h1>{React.string("here")}</h1>
        
        </div>
    </form>
};