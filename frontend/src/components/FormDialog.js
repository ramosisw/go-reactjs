import React from 'react';
import PropTypes from 'prop-types';
import Button from '@material-ui/core/Button';
import { withStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogTitle from '@material-ui/core/DialogTitle';

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
            autoFocus
            margin="dense"
            id="titte"
            label="Title"
            type="text"
            fullWidth
          />
          <TextField
            value={todo.description}
            multiline
            margin="dense"
            id="name"
            label="Description"
            type="text"
            // onChange={handleChange('multiline')}
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


FormDialog.propTypes = {
  open: PropTypes.bool.isRequired,
  onClose: PropTypes.func.isRequired,
  todo: PropTypes.object.isRequired
};

export default withStyles(styles)(FormDialog);