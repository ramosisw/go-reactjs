import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import Button from '@material-ui/core/Button';
import { withStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogTitle from '@material-ui/core/DialogTitle';
import { todoActions } from '../redux/todo/todo.actions';

const styles = theme => ({

});

class FormDialog extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      todo: props.todo
    }
  }

  onSave = () => {
    const { todo } = this.state;
    const { edit, onClose } = this.props
    if (edit) {
      this.props.dispatch(todoActions.put(todo));
    } else {
      this.props.dispatch(todoActions.post(todo));
    }
    onClose();
  }

  handleChange = (e) => {
    const { name, value } = e.target;
    const { todo } = this.state;
    this.setState({ todo: { ...todo, [name]: value } });
  }

  render() {
    const { open, onClose } = this.props
    const { todo } = this.state;
    return (
      <Dialog open={open} onClose={onClose} aria-labelledby="form-dialog-title">
        <DialogTitle id="form-dialog-title">Fill fields</DialogTitle>
        <DialogContent>
          <TextField
            value={todo.title}
            onChange={this.handleChange}
            autoFocus
            autoComplete={"off"}
            margin="dense"
            name="title"
            label="Title"
            type="text"
            fullWidth
          />
          <TextField
            value={todo.description}
            onChange={this.handleChange}
            multiline
            margin="dense"
            name="description"
            label="Description"
            type="text"
            rows={4}
            fullWidth
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={onClose} color="primary" >Cancel</Button>
          <Button onClick={this.onSave} >Save</Button>
        </DialogActions>
      </Dialog>
    )
  }
}

function mapStateToProps(state) {
  const { data } = state.todo.get;
  return {
    data
  };
}


FormDialog.propTypes = {
  open: PropTypes.bool.isRequired,
  onClose: PropTypes.func.isRequired,
  todo: PropTypes.object.isRequired,
  edit: PropTypes.bool.isRequired,
};

export default withStyles(styles)(connect(mapStateToProps)(FormDialog));