const Header = () => (
    <Fragment>
        <h1>Example</h1>
        <h4>for Preact 10</h4>
    </Fragment>
);

const Item = ({ message }) => (
    <li>{message}</li>
);

const List = ({ items }) => {
    return (
        <ul>
            { items.map((message) => (
                <Item message={message} />
            ))}
        </ul>
    )
}

const { Items } = data;

render(
    <App>
        <Header />
        <List items={Items} />
    </App>
)
