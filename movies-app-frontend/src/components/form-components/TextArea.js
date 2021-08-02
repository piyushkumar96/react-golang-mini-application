const TextArea = (props) => {
  return (
    <div className="mb-3">
      <label htmlFor="description" className="form-label">
        {props.title}
      </label>
      <textarea
        className={`form-control ${props.className}`}
        id={props.name}
        name={props.name}
        value={props.value}
        onChange={props.handleChange}
        rows={props.rows}
      />
      <div className={props.errorDiv}>{props.errorMsg}</div>
    </div>
  );
};

export default TextArea;
