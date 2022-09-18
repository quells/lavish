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
        <List items={Items} />
    </App>
)
