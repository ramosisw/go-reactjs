import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Moment from 'react-moment';
import Card from '@material-ui/core/Card';
import CardHeader from '@material-ui/core/CardHeader';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardContent from '@material-ui/core/CardContent';
import Typography from '@material-ui/core/Typography';

const styles = theme => ({
    card: {
        // minWidth: 200,
        // maxWidth: 345,
        marginBottom: theme.spacing(1),
    },
    subheader: {
        textAlign: 'right',
        display: 'block',
        fontSize: 14
    }

});

function AppCard(props) {
    const { classes, todo } = props;

    return (
        <Card className={classes.card}>
            <CardActionArea>
                <CardHeader
                    title={todo.title}
                    subheader={<Moment format="ddd DD MMM 'YY" className={classes.subheader}>{todo.CreatedAt}</Moment>}
                />
                <CardContent>
                    <Typography component="p">
                        {todo.description}
                    </Typography>
                </CardContent>
            </CardActionArea>
            {/* <CardActions>
                <Button size="small" color="primary">
                    Done
                </Button>
                <Button size="small" color="primary">
                    Edit
                </Button>
            </CardActions> */}
        </Card >
    );
}

AppCard.propTypes = {
    classes: PropTypes.object.isRequired
};

export default withStyles(styles)(AppCard);
